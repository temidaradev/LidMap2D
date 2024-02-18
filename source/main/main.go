package game

import (
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lidldev/LidMap2D/source/menu"
)

type Game struct {
	m     menu.Menu
	mFull menu.MenuFullScreen
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
		g.mFull.TitleFullScreen(screen)
		g.mFull.CreateButtonFullScreen(screen)
		g.mFull.SettingButtonFullScreen(screen)
	} else {
		g.m.Title(screen)
		g.m.CreateButton(screen)
		g.m.SettingButton(screen)
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