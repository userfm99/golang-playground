package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/test", upload)

	e.Start(":8080")
}

func upload(c echo.Context) error {
	// Source
	file, err := c.FormFile("field")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	var img image.Image
	if strings.Contains(file.Filename, "jpg") || strings.Contains(file.Filename, "jpeg") {
		img, err = jpeg.Decode(src)
		if err != nil {
			log.Fatal(err)
		}
	}
	if strings.Contains(file.Filename, "png") {
		img, err = png.Decode(src)
		if err != nil {
			log.Fatal(err)
		}
	}

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(1200, 370, img, resize.Lanczos3)

	out, err := os.Create("new" + file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, out); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully</p>", file.Filename))
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}