package game

import (
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	createMap     bool
	createFullMap bool
	m             Menu
	mFull         MenuFullScreen
}

func NewGame() *Game {
	g := &Game{}

	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if ebiten.IsFullscreen() {
		// g.mFull.TitleFullScreen(screen)
		// g.mFull.CreateButtonFullScreen(screen, g)
		// g.mFull.SettingButtonFullScreen(screen)
		g.mFull.showFullMenu(screen, g)
		if g.createFullMap {
			screen.Fill(color.White)
		}
	} else {
		// g.m.Title(screen)
		// g.m.CreateButton(screen, g)
		// g.m.SettingButton(screen)

		g.m.showMenu(screen, g)
	}

	if g.createMap {
		screen.Fill(color.Black)
		g.m.clearMenu(screen)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		f, err := os.Create("screenshot.png")
		if err != nil {
			log.Fatal("can't create file: ", err)
		}
		png.Encode(f, screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
