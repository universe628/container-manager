package handler

import (
	"container-manager/internal/schema"
	"context"
	"mime/multipart"
)

type UserService interface {
}

type AuthService interface {
	Login(ctx context.Context, req *schema.User) (string, error)
	NewUser(ctx context.Context, req *schema.User) error
}

type FileService interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader) error
}

type Validationservice interface {
	FileUploadValidation(files []*multipart.FileHeader) error
}
