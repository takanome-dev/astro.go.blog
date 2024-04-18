package utils

import (
	"context"
	"mime/multipart"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func CloudinaryUpload(file multipart.File, filename string) (*uploader.UploadResult, error) {
	cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	
	resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{PublicID: strings.Split(filename, ".")[0], Folder: "astro-blog"});
	
	return resp, err
}