package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/w0termel0n/SilyHorrorGame/silygame"
)

func main() {
	g := silygame.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
