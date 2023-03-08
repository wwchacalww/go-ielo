package utils

import (
	"fmt"
	"mime/multipart"

	"github.com/disintegration/imaging"
)

func AvatarIsValid(fh *multipart.FileHeader) error {
	var maxSize int64
	maxSize = 15 * 1024 * 1024 // 15mb
	ct := fh.Header.Get("Content-Type")

	if ct != "image/jpeg" {
		return fmt.Errorf("File type invalid")
	}

	if fh.Size > maxSize {
		return fmt.Errorf("Avatar file must be smaller than 15mb")
	}

	return nil
}

func AvatarResize(filename string) error {
	src, err := imaging.Open("public/imgs/" + filename)
	if err != nil {
		return err
	}
	if src.Bounds().Dx() > src.Bounds().Dy() {
		src = imaging.Resize(src, 0, 400, imaging.Lanczos)
	} else {
		src = imaging.Resize(src, 400, 0, imaging.Lanczos)
	}
	src = imaging.Fill(src, 400, 400, imaging.Center, imaging.Lanczos)
	err = imaging.Save(src, "public/imgs/"+filename)
	return err
}
