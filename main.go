package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/w0termel0n/SilyHorrorGame/silygame"
)

func (g *Game) Update() error {
	return nil
}

func (g *silygame.Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := silygame.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
