package filerepo

import (
	errs "container-manager/internal/errors"
	"container-manager/internal/schema"
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
)

type FileRepo struct {
	fileManager FileManager
}

func NewLocalStorageRepo(fileManager FileManager) *FileRepo {
	return &FileRepo{
		fileManager: fileManager,
	}
}

func (r *FileRepo) SaveFile(ctx context.Context, file *multipart.FileHeader) error {
	relativePath := filepath.Clean(file.Filename)
	fmt.Println(relativePath)
	userID, ok := ctx.Value(schema.UserIDKey).(string)
	if !ok {
		return errs.ErrUnknownUser
	}
	path := filepath.Join("uploads", userID, relativePath)

	return r.fileManager.WriteFile(ctx, file, path)
}

func (r *FileRepo) DownloadFile() {

}
