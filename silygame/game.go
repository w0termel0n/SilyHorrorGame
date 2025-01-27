package silygame

import (
	//"fmt"
	//"image/color"
	"time"
	//"github.com/hajimehoshi/ebiten/v2"
	//"github.com/w0termel0n/SilyHorrorGame/assets"
)

const (
	screen = 800
	width  = 600

	baseVelocity = 0.25
	duration     = 15 * time.Second
)

type Game struct {
	//player       *Player
	velocity float64
	velTimer *Timer
}

func NewGame() *Game {
	g := &Game{
		velocity: baseVelocity,
		velTimer: NewTimer(duration),
	}

	//g.player = NewPlayer(g)

	return g
}

/*
func (g *Game) Update() error {

}

func (g *Game) Draw(screen ebiten.Image) {

}
*/
