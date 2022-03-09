package dxo

import (
	"encoding/json"

	"github.com/bitwormhole/starter/util"
)

////////////////////////////////////////////////////////////////////////////////

// FileManifest 对象表示一个包含多个文件的清单
type FileManifest struct {
	Version int                 `json:"version"` // current=1
	Items   []*FileManifestItem `json:"items"`
}

////////////////////////////////////////////////////////////////////////////////

// FileManifestItem 表示清单中的一个条目（即一个文件）
type FileManifestItem struct {
	Name        string `json:"name"` // 文件名
	ContentType string `json:"contenttype"`
	URL         string `json:"url"`
	Sha256sum   string `json:"sha256sum"`
	Size        int64  `json:"size"`
}

////////////////////////////////////////////////////////////////////////////////

// FileManifestValue 是 FileManifest 对象经过编码的形式
type FileManifestValue string

// GetManifest 获取清单对象
func (v FileManifestValue) GetManifest() *FileManifest {
	str := v.String()
	b64 := util.Base64FromString(str)
	bin := b64.Bytes()
	o := &FileManifest{}
	json.Unmarshal(bin, o)
	return o
}

func (v FileManifestValue) String() string {
	return string(v)
}

// ValueOfFileManifest 把 FileManifest 对象转换成base64值
func ValueOfFileManifest(fm *FileManifest) FileManifestValue {
	bin, err := json.Marshal(fm)
	if err != nil {
		return "{}"
	}
	b64 := util.Base64FromBytes(bin)
	v := b64.String()
	return FileManifestValue(v)
}

////////////////////////////////////////////////////////////////////////////////
