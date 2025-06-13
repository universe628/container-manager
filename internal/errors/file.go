package errs

import "errors"

var ErrNoFileUploaded = errors.New("no file uploaded")
var ErrFileTooLarge = errors.New("exceeded maximum file size limit")
var ErrFileRead = errors.New("error reading file")
var ErrInvalidFileType = errors.New("invalid file type")
var ErrInvalidFilePath = errors.New("invalid file path")
