package game

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	} else {
		g.m.Title(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
