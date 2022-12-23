package api

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
)

func ConvertImageToBuffer(new_image image.Image) *bytes.Reader {
	// create buffer
	buff := new(bytes.Buffer)

	// encode image to buffer
	err := png.Encode(buff, new_image)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}

	// convert buffer to reader
	reader := bytes.NewReader(buff.Bytes())

	return reader
}
