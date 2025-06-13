package file

import (
	"mime/multipart"
	"net/http"
)

func CheckMime(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	file.Seek(0, 0)

	mimeType := http.DetectContentType(buffer)
	return mimeType, nil
}
