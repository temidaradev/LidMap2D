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
		bigFontSize    = 100
	)

	const x = 490

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, 140)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, "LidMap2D", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   bigFontSize,
	}, op)
}

func (m *MenuFullScreen) CreateButtonFullScreen(screen *ebiten.Image) {
	button := ebiten.NewImage(180, 50)
	button.Fill(color.RGBA{50, 50, 50, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(630, 400)

	screen.DrawImage(button, op)

	op2 := &text.DrawOptions{}
	op2.GeoM.Translate(683, 409)
	op2.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, "Create", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   normalFontSize,
	}, op2)
}

func (m *MenuFullScreen) SettingButtonFullScreen(screen *ebiten.Image) {
	button := ebiten.NewImage(180, 50)
	button.Fill(color.RGBA{50, 50, 50, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(630, 465)

	screen.DrawImage(button, op)

	op2 := &text.DrawOptions{}
	op2.GeoM.Translate(674, 473)
	op2.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, "Settings", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   normalFontSize,
	}, op2)
}
