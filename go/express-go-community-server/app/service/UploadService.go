package service

import (
	"context"

	"github.com/bitwomrhole-express/express/community-server/app/web/dto"
)

// UploadService 提供上传文件的服务
// 【inject:"#express-UploadService"】
type UploadService interface {
	BeginUpload(c context.Context, up *dto.Upload) (*dto.Upload, error)
	EndUpload(c context.Context, id string, up *dto.Upload) (*dto.Upload, error)
}
