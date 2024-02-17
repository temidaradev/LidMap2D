package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	game "github.com/lidldev/LidMap2D/main"
)

func main() {
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
