package silygame

/*
import (
	"fmt"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)
*/
const (
	screen = 800
	width  = 600
)

type Game struct {
	player       *Player
	baseVelocity float64
	velTimer     *Timer
}
