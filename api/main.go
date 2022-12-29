package api

import (
	"bytes"
	"cutiecat6778/discordbot/utils"
	"image"
	"image/png"
)

func ConvertImageToBuffer(newImage image.Image) *bytes.Reader {
	// create buffer
	buff := new(bytes.Buffer)

	// encode image to buffer
	err := png.Encode(buff, newImage)
	if err != nil {
		utils.HandleServerError(err)
	}

	// convert buffer to reader
	reader := bytes.NewReader(buff.Bytes())

	return reader
}
