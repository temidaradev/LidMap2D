package game

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/lidldev/LidMap2D/assets"
)

type MenuFullScreen struct {
	createMap     bool
	createFull    *ebiten.Image
	createSettigs *ebiten.Image
}

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(assets.Font_ttf))
	if err != nil {
		log.Fatal(err)
	}
	fontFaceSource = s
}

func (m *MenuFullScreen) DrawFull(screen *ebiten.Image) {
	m.TitleFullScreen(screen)
	m.CreateButtonFullScreen(screen, &Game{})
	m.SettingButtonFullScreen(screen)
}

func (m *MenuFullScreen) UpdateFull(g *Game) error {
	cx, cy := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if m.createFull.Bounds().Min.X+630 <= cx && cx < m.createFull.Bounds().Max.X+630 && m.createFull.Bounds().Min.Y+400 <= cy && cy < m.createFull.Bounds().Max.Y+400 {
			log.Print("dflkbnkfbn")
			g.md = true
		}
	}
	return nil
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

func (m *MenuFullScreen) CreateButtonFullScreen(screen *ebiten.Image, g *Game) {
	m.createFull = ebiten.NewImage(180, 50)
	m.createFull.Fill(color.RGBA{50, 50, 50, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(630, 400)

	screen.DrawImage(m.createFull, op)

	op2 := &text.DrawOptions{}
	op2.GeoM.Translate(683, 409)
	op2.ColorScale.ScaleWithColor(color.White)

	cx, cy := ebiten.CursorPosition()

	if m.createFull.Bounds().Min.X+630 <= cx && cx < m.createFull.Bounds().Max.X+630 && m.createFull.Bounds().Min.Y+400 <= cy && cy < m.createFull.Bounds().Max.Y+400 {
		op2.ColorScale.ScaleWithColor(color.RGBA{20, 20, 30, 255})
		m.createFull.Fill(color.RGBA{70, 70, 70, 255})
		screen.DrawImage(m.createFull, op)
	}

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

	cx, cy := ebiten.CursorPosition()

	if button.Bounds().Min.X+630 <= cx && cx < button.Bounds().Max.X+630 && button.Bounds().Min.Y+465 <= cy && cy < button.Bounds().Max.Y+465 {
		op2.ColorScale.ScaleWithColor(color.RGBA{20, 20, 30, 255})
		button.Fill(color.RGBA{70, 70, 70, 255})
		screen.DrawImage(button, op)

		op3 := &text.DrawOptions{}
		op3.GeoM.Translate(583, 530)
		op3.ColorScale.ScaleWithColor(color.White)

		text.Draw(screen, "Coming Not Soon!", &text.GoTextFace{
			Source: fontFaceSource2,
			Size:   midFontSize,
		}, op3)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if button.Bounds().Min.X+630 <= cx && cx < button.Bounds().Max.X+630 && button.Bounds().Min.Y+465 <= cy && cy < button.Bounds().Max.Y+465 {
			log.Printf("Settings")
		}
	}

	text.Draw(screen, "Settings", &text.GoTextFace{
		Source: fontFaceSource,
		Size:   normalFontSize,
	}, op2)
}

func (m *MenuFullScreen) showFullMenu(screen *ebiten.Image, g *Game) {
	m.TitleFullScreen(screen)
	m.CreateButtonFullScreen(screen, g)
	m.SettingButtonFullScreen(screen)
}
