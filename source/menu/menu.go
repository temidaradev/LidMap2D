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

type Menu struct{}

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
	button := ebiten.NewImage(150, 40)
	button.Fill(color.RGBA{50, 50, 50, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(240, 200)

	screen.DrawImage(button, op)

	op2 := &text.DrawOptions{}
	op2.GeoM.Translate(278, 204)
	op2.ColorScale.ScaleWithColor(color.White)

	cx, cy := ebiten.CursorPosition()

	if button.Bounds().Min.X+240 <= cx && cx < button.Bounds().Max.X+240 && button.Bounds().Min.Y+200 <= cy && cy < button.Bounds().Max.Y+200 {
		op2.ColorScale.ScaleWithColor(color.RGBA{20, 20, 30, 255})
		button.Fill(color.RGBA{70, 70, 70, 255})
		screen.DrawImage(button, op)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if button.Bounds().Min.X+240 <= cx && cx < button.Bounds().Max.X+240 && button.Bounds().Min.Y+200 <= cy && cy < button.Bounds().Max.Y+200 {
			log.Printf("Create")
		}
	}

	text.Draw(screen, "Create", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   normalFontSize,
	}, op2)
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

	cx, cy := ebiten.CursorPosition()

	if button.Bounds().Min.X+240 <= cx && cx < button.Bounds().Max.X+240 && button.Bounds().Min.Y+250 <= cy && cy < button.Bounds().Max.Y+250 {
		op2.ColorScale.ScaleWithColor(color.RGBA{20, 20, 30, 255})
		button.Fill(color.RGBA{70, 70, 70, 255})
		screen.DrawImage(button, op)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if button.Bounds().Min.X+240 <= cx && cx < button.Bounds().Max.X+240 && button.Bounds().Min.Y+250 <= cy && cy < button.Bounds().Max.Y+250 {
			log.Printf("Settings")
		}
	}

	text.Draw(screen, "Settings", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   normalFontSize,
	}, op2)
}
