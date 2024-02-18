package menu

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/lidldev/LidMap2D/assets"
)

type Menu struct {
}

var (
	fontFaceSource *text.GoTextFaceSource
)

var (
	Grey = color.RGBA{25, 25, 25, 255}
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(assets.Font_ttf))
	if err != nil {
		log.Fatal(err)
	}
	fontFaceSource = s
}

func (m *Menu) Title(screen *ebiten.Image) {
	screen.Fill(Grey)
	const (
		normalFontSize = 24
		bigFontSize    = 48
	)

	const x = 200

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, 60)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, "LidMap2D", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   bigFontSize,
	}, op)
}

func (m *Menu) Button() {

}
