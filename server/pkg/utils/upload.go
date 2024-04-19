package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func HandleImage(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	defer file.Close()

	fileSize := fileHeader.Size

	if fileSize > MAX_UPLOAD_SIZE {
		return "", fmt.Errorf("file is too big, the maximum allowed is 2MB")
	}

	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		return "", err
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		return "", fmt.Errorf("the provided file format is not allowed. Please upload a JPEG or PNG image")
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	resp, err := CloudinaryUpload(file, strings.Split(fileHeader.Filename, ".")[0])
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}