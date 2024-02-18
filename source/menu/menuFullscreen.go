package menu

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/lidldev/LidMap2D/assets"
)

type MenuFullScreen struct{}

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(assets.Font_ttf))
	if err != nil {
		log.Fatal(err)
	}
	fontFaceSource = s
}

func (m *MenuFullScreen) TitleFullScreen(screen *ebiten.Image) {
	screen.Fill(Grey)
	const (
		normalFontSize = 24
		bigFontSize    = 64
	)

	const x = 570

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, 140)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, "LidMap2D", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   bigFontSize,
	}, op)
}
