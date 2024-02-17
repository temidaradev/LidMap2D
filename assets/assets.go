package assets

import (
	"embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

//go:embed fonts/AvenirNext-DemiBoldItalic.ttf
var Font_ttf []byte

func GetSingleImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
