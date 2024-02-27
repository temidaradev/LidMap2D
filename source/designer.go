package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Designer struct {
}

func (d *Designer) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
}

func (d *Designer) Update() error {
	return nil
}
