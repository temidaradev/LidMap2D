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
	createMap        bool
	createFullMap    bool
	isButtonDeactive bool
	m                Menu
	mFull            MenuFullScreen
}

func NewGame() *Game {
	g := &Game{}

	return g
}

func (g *Game) Update() error {
	cx, cy := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.m.createButton.Bounds().Min.X+240 <= cx && cx < g.m.createButton.Bounds().Max.X+240 && g.m.createButton.Bounds().Min.Y+200 <= cy && cy < g.m.createButton.Bounds().Max.Y+200 {
			log.Print("dflkbnkfbn")
			g.createMap = true
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if ebiten.IsFullscreen() {
		g.mFull.showFullMenu(screen, g)
		if g.createFullMap {
			screen.Fill(color.White)
		}
	} else {
		g.m.showMenu(screen, g)
	}

	if g.createMap {
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
