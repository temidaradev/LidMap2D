package game

import (
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	m            Menu
	d            Designer
	md           bool
	isFullScreen bool
	mFull        MenuFullScreen
}

func NewGame() *Game {
	g := &Game{}

	return g
}

func (g *Game) Update() error {
	full := ebiten.IsFullscreen()
	if !g.md {
		g.m.Update(g)
		if full {
			g.mFull.UpdateFull(g)
		}
	} else {
		g.d.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	full := ebiten.IsFullscreen()
	if !g.md {
		g.m.Draw(screen)
		if full {
			g.mFull.DrawFull(screen)
		}
	} else {
		g.d.Draw(screen)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		f, err := os.Create("screenshot.png")
		if err != nil {
			log.Fatal("can't create file: ", err)
		}
		log.Print("Got A Screenshot")
		png.Encode(f, screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
