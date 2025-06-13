package fileservice

import (
	"context"
	"mime/multipart"
)

type FileRepo interface {
	SaveFile(ctx context.Context, file *multipart.FileHeader) error
}
