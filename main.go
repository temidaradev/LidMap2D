package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	game "github.com/lidldev/LidMap2D/source"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	ebiten.SetWindowTitle("LidMap2D")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

//knasdghfpouh
