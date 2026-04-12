package _115_open

import (
	"github.com/OpenListTeam/OpenList/v4/internal/driver"
	"github.com/OpenListTeam/OpenList/v4/internal/op"
)

type Addition struct {
	// Usually one of two
	driver.RootID
	// define other
	OrderBy        string  `json:"order_by" type:"select" options:"file_name,file_size,user_utime,file_type"`
	OrderDirection string  `json:"order_direction" type:"select" options:"asc,desc"`
	LimitRate      float64 `json:"limit_rate" type:"float" default:"1" help:"limit all api request rate ([limit]r/1s)"`
	PageSize       int64   `json:"page_size" type:"number" default:"200" help:"list api per page size of 115open driver"`
	AccessToken    string  `json:"access_token" required:"true"`
	RefreshToken   string  `json:"refresh_token" required:"true"`
	// 115 逆向驱动配置（用于获取缩略图）
	Cookie115       string `json:"cookie_115" type:"text" help:"115 cookie for thumbnail"`
	QRCodeToken115  string `json:"qrcode_token_115" type:"text" help:"115 QR code token for thumbnail"`
	QRCodeSource115 string `json:"qrcode_source_115" type:"select" options:"web,android,ios,tv,alipaymini,wechatmini,qandroid" default:"linux" help:"select the QR code device for thumbnail"`
}

var config = driver.Config{
	Name:          "115 Open",
	DefaultRoot:   "0",
	LinkCacheMode: driver.LinkCacheUA,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &Open115{}
	})
}
