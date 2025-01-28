// the files 'go.mod' and 'go.sum' keep up with your dependencies and stuff
//
// 'go.mod' is required to run a go program because it also holds the path to
// the designated module to run

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
