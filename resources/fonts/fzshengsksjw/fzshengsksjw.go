package fzshengsksjw

import (
	assets "github.com/cmict-aict/go-captcha-assets/bindata/fonts/fzshengsksjw"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func GetFont() (font *truetype.Font, err error) {
	asset, err := assets.Asset("sourcedata/fonts/fzshengsksjw_cu.ttf")
	if err != nil {
		return font, err
	}

	font, err = freetype.ParseFont(asset)
	if err != nil {
		return nil, err
	}

	return font, nil
}
