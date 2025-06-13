package fileservice

import (
	"context"
	"mime/multipart"
)

type fileService struct {
	fileRepo FileRepo
}

func NewFileService(fileRepo FileRepo) *fileService {
	return &fileService{
		fileRepo: fileRepo,
	}
}

func (s *fileService) UploadFile(gtx context.Context, file *multipart.FileHeader) error {
	return s.fileRepo.SaveFile(gtx, file)
}

