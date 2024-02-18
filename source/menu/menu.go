package menu

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/lidldev/LidMap2D/assets"
)

type Menu struct {
	button *ebiten.Image
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

const (
	normalFontSize = 24
	bigFontSize    = 48
)

func (m *Menu) Title(screen *ebiten.Image) {
	screen.Fill(Grey)

	const x = 200

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, 60)
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, "LidMap2D", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   bigFontSize,
	}, op)
}

func (m *Menu) CreateButton(screen *ebiten.Image) {
	m.button = ebiten.NewImage(150, 40)
	m.button.Fill(color.RGBA{50, 50, 50, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(240, 200)

	screen.DrawImage(m.button, op)

	op2 := &text.DrawOptions{}
	op2.GeoM.Translate(278, 204)
	op2.ColorScale.ScaleWithColor(color.White)

	cx, cy := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if m.button.Bounds().Min.X+240 <= cx && cx < m.button.Bounds().Max.X+240 && m.button.Bounds().Min.Y+200 <= cy && cy < m.button.Bounds().Max.Y+200 {
			op2.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
			log.Printf("nbnbnbnbnbnbnn")
		}
	}

	text.Draw(screen, "Create", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   normalFontSize,
	}, op2)

	//if image.Pt(x, y).In(rect.Bounds().Add(translateX, translateY)) {
}

func (m *Menu) CreateButtonAction() {
	// cx, cy := ebiten.CursorPosition()
	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	// 	if m.button.Bounds().Min.X+240 <= cx && cx < m.button.Bounds().Max.X+240 && m.button.Bounds().Min.Y+200 <= cy && cy < m.button.Bounds().Max.Y+200 {
	// 		log.Print("nngggsdgsd")
	// 	}
	// }
}

func (m *Menu) SettingButton(screen *ebiten.Image) {
	button := ebiten.NewImage(150, 40)
	button.Fill(color.RGBA{50, 50, 50, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(240, 250)

	screen.DrawImage(button, op)

	op2 := &text.DrawOptions{}
	op2.GeoM.Translate(270, 254)
	op2.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, "Settings", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   normalFontSize,
	}, op2)
}
