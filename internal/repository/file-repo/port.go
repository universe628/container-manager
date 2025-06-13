package filerepo

import (
	"context"
	"mime/multipart"
	"os"
)

type FileManager interface {
	ReadFile(ctx context.Context, path string) ([]byte, error)
	ReadDirectory(ctx context.Context, path string) ([]os.DirEntry, error)
	WriteFile(ctx context.Context, file *multipart.FileHeader, path string) error
	DeleteFile(ctx context.Context, path string) error
}
