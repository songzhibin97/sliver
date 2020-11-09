package interfaces

import (
	"github.com/SliverHorn/sliver/global"
	"github.com/SliverHorn/sliver/integration/oss"
	"mime/multipart"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func NewOss() OSS {
	switch global.Config.System.OssType {
	case "local":
		return &oss.Local{}
	case "qiniu":
		return &oss.Qiniu{}
	default:
		return &oss.Local{}
	}
}