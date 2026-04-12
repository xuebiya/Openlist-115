package _115_open

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	"github.com/OpenListTeam/OpenList/v4/internal/conf"
	driver115 "github.com/SheltonZhu/115driver/pkg/driver"
	"github.com/pkg/errors"
)

func ParseInt64(v json.Number) (int64, error) {
	i, err := v.Int64()
	if err == nil {
		return i, nil
	}
	f, e1 := v.Float64()
	if e1 == nil {
		return int64(f), nil
	}
	return int64(0), err
}

// ThumbClient 115 逆向客户端，用于获取缩略图
type ThumbClient struct {
	client *driver115.Pan115Client
}

var thumbClient *ThumbClient

// getThumbUA 获取缩略图客户端的 UserAgent
func getThumbUA() string {
	return fmt.Sprintf("Mozilla/5.0 115Browser/%s", appVer)
}

// InitThumbClient 初始化 115 缩略图客户端
func InitThumbClient(cookie, qrcodeToken, qrcodeSource string) error {
	// 初始化 appVer
	initAppVer()

	// 如果都没有配置，禁用缩略图
	if cookie == "" && qrcodeToken == "" {
		thumbClient = nil
		return nil
	}

	opts := []driver115.Option{
		driver115.UA(getThumbUA()),
		func(c *driver115.Pan115Client) {
			c.Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: conf.Conf.TlsInsecureSkipVerify})
		},
	}

	client := driver115.New(opts...)
	cr := &driver115.Credential{}

	if qrcodeToken != "" {
		// QRCode 登录
		s := &driver115.QRCodeSession{
			UID: qrcodeToken,
		}
		var err error
		cr, err = client.QRCodeLoginWithApp(s, driver115.LoginApp(qrcodeSource))
		if err != nil {
			return errors.Wrap(err, "failed to login by QR code")
		}
	} else if cookie != "" {
		// Cookie 登录
		var err error
		if err = cr.FromCookie(cookie); err != nil {
			return errors.Wrap(err, "failed to parse cookie")
		}
		client.ImportCredential(cr)
	} else {
		return errors.New("missing cookie or QR code token")
	}

	if err := client.LoginCheck(); err != nil {
		return errors.Wrap(err, "failed to login check")
	}

	thumbClient = &ThumbClient{
		client: client,
	}
	return nil
}

// GetThumbsByIDs 批量获取缩略图
func (c *ThumbClient) GetThumbsByIDs(fileIDs []string) (map[string]string, error) {
	result := make(map[string]string)

	// 使用 GetFile API 逐个获取文件信息
	for _, fileID := range fileIDs {
		file, err := c.client.GetFile(fileID)
		if err != nil {
			continue
		}
		if file.ThumbURL != "" {
			result[fileID] = file.ThumbURL
		}
	}

	return result, nil
}
