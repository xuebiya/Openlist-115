package _115_open

import (
	"time"

	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/OpenListTeam/OpenList/v4/pkg/utils"
	sdk "github.com/OpenListTeam/115-sdk-go"
)

type Obj sdk.GetFilesResp_File

// Thumb implements model.Thumb.
func (o *Obj) Thumb() string {
	// 如果有 115 逆向 cookie，缩略图会通过 Thumbnail115 获取
	return o.Thumbnail
}

// ObjWithThumb 扩展 Obj，添加 115 缩略图获取支持
type ObjWithThumb struct {
	*Obj
	ThumbURL string
}

// Thumb115 获取 115 缩略图
func (o *ObjWithThumb) Thumb() string {
	return o.ThumbURL
}

// CreateTime implements model.Obj.
func (o *Obj) CreateTime() time.Time {
	return time.Unix(o.UpPt, 0)
}

// CreateTime implements model.Obj.
func (o *ObjWithThumb) CreateTime() time.Time {
	return time.Unix(o.Obj.UpPt, 0)
}

// GetHash implements model.Obj.
func (o *Obj) GetHash() utils.HashInfo {
	return utils.NewHashInfo(utils.SHA1, o.Sha1)
}

// GetHash implements model.Obj.
func (o *ObjWithThumb) GetHash() utils.HashInfo {
	return utils.NewHashInfo(utils.SHA1, o.Obj.Sha1)
}

// GetID implements model.Obj.
func (o *Obj) GetID() string {
	return o.Fid
}

// GetID implements model.Obj.
func (o *ObjWithThumb) GetID() string {
	return o.Obj.Fid
}

// GetName implements model.Obj.
func (o *Obj) GetName() string {
	return o.Fn
}

// GetName implements model.Obj.
func (o *ObjWithThumb) GetName() string {
	return o.Obj.Fn
}

// GetPath implements model.Obj.
func (o *Obj) GetPath() string {
	return ""
}

// GetPath implements model.Obj.
func (o *ObjWithThumb) GetPath() string {
	return ""
}

// GetSize implements model.Obj.
func (o *Obj) GetSize() int64 {
	return o.FS
}

// GetSize implements model.Obj.
func (o *ObjWithThumb) GetSize() int64 {
	return o.Obj.FS
}

// IsDir implements model.Obj.
func (o *Obj) IsDir() bool {
	return o.Fc == "0"
}

// IsDir implements model.Obj.
func (o *ObjWithThumb) IsDir() bool {
	return o.Obj.Fc == "0"
}

// ModTime implements model.Obj.
func (o *Obj) ModTime() time.Time {
	return time.Unix(o.Upt, 0)
}

// ModTime implements model.Obj.
func (o *ObjWithThumb) ModTime() time.Time {
	return time.Unix(o.Obj.Upt, 0)
}

var _ model.Obj = (*Obj)(nil)
var _ model.Thumb = (*Obj)(nil)
var _ model.Obj = (*ObjWithThumb)(nil)
var _ model.Thumb = (*ObjWithThumb)(nil)
