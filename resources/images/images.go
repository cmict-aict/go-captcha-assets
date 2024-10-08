package images

import (
	"fmt"
	"image"

	assets "github.com/cmict-aict/go-captcha-assets/bindata/images"
	"github.com/cmict-aict/go-captcha-assets/helper"
)

func GetImages() ([]image.Image, error) {
	var images []image.Image
	var err error
	for i := 1; i <= 10; i++ {
		var img image.Image
		var asset []byte
		asset, err = assets.Asset(fmt.Sprintf("sourcedata/images/image-%d.png", i))
		if err != nil {
			return images, err
		}

		img, err = helper.DecodeByteToJpeg(asset)
		if err != nil {
			return images, err
		}

		images = append(images, img)
	}

	return images, nil
}
