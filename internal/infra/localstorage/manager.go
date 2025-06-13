package localstorage

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type fileManager struct{}

func NewfileManager() *fileManager {
	return &fileManager{}
}

func (fm *fileManager) CreateDirectory(ctx context.Context, path string) error {
	return os.MkdirAll(path, 0750)
}

func (fm *fileManager) ReadFile(ctx context.Context, path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (fm *fileManager) ReadDirectory(ctx context.Context, path string) ([]os.DirEntry, error) {
	return os.ReadDir(path)
}

func (fm *fileManager) WriteFile(ctx context.Context, file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = fm.CreateDirectory(ctx, filepath.Dir(dst)); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (fm *fileManager) DeleteFile(ctx context.Context, path string) error {
	return os.Remove(path)
}
