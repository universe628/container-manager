package validationservice

import (
	errs "container-manager/internal/errors"
	"container-manager/internal/infra/config"
	"container-manager/internal/infra/file"
	"mime/multipart"
	"path/filepath"
	"strings"
)

type validationService struct{}

func NewValidationService() *validationService {
	return &validationService{}
}

const (
	KB = 1024
	MB = 1024 * KB
)

var allowedMIIMETypes = map[string]bool{
	"text/csv":   true,
	"image/jpeg": true,
	"image/png":  true,
}

func isAllowedFileType(fileType string) bool {
	return allowedMIIMETypes[fileType]
}

func (v *validationService) FileUploadValidation(files []*multipart.FileHeader) error {
	if len(files) == 0 {
		return errs.ErrNoFileUploaded
	}

	for i := 0; i < len(files); i++ {
		if files[i].Size > int64(config.GetConfig().MaximunFileSizeMB)*MB {
			return errs.ErrFileTooLarge
		}

		cleanRelPath := filepath.Clean(files[i].Filename)
		if strings.Contains(cleanRelPath, "..") {
			return errs.ErrInvalidFilePath
		}

		fileContent, err := files[i].Open()
		if err != nil {
			return errs.ErrFileRead
		}
		defer fileContent.Close()

		fileType, err := file.CheckMime(fileContent)
		if err != nil {
			return errs.ErrFileRead
		}
		if !isAllowedFileType(fileType) {
			return errs.ErrInvalidFileType
		}
	}
	return nil
}
