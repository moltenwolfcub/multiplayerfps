package assets

import (
	"bytes"
	"image"
	_ "image/png"
)

func LoadPNG(file string) (image.Image, error) {

	embeddedImage, err := textures.ReadFile("textures/" + file + ".png")
	if err != nil {
		return nil, err
	}

	image, _, err := image.Decode(bytes.NewReader(embeddedImage))
	if err != nil {
		return nil, err
	}
	return image, nil
}

func MustLoadPNG(file string) image.Image {
	img, err := LoadPNG(file)
	if err != nil {
		panic("Failed to load PNG: " + err.Error())
	}
	return img
}

var (
	Metal_full image.Image = MustLoadPNG("metal/metalbox_full")
)
