package dto

import (
	"mime/multipart"
)

type UploadFilesInput struct {
	Files []*multipart.FileHeader
}
