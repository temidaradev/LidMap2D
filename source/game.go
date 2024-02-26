package game

import (
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	m Menu
	//mFull            MenuFullScreen
}

func NewGame() *Game {
	g := &Game{}

	return g
}

func (g *Game) Update() error {
	g.m.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.m.Draw(screen)

	if g.m.isCreateButton {
		screen.Clear()
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
