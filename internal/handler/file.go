package handler

import (
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileService       FileService
	validationservice Validationservice
}

func NewFileHandler(fileService FileService, validationService Validationservice) *FileHandler {
	return &FileHandler{
		fileService:       fileService,
		validationservice: validationService,
	}
}

func (h *FileHandler) UploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.Error(err)
		return
	}

	err = h.validationservice.FileUploadValidation(form.File["file"])
	if err != nil {
		c.Error(err)
		return
	}

	for i := 0; i < len(form.File["file"]); i++ {
		err = h.fileService.UploadFile(c.Request.Context(), form.File["file"][i])
		if err != nil {
			c.Error(err)
			return
		}
	}

}
