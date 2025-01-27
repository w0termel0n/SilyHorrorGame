package silygame

const (
	screen = 800
	width  = 600
)

type Game struct {
	player       *Player
	baseVelocity float64
	velTimer     *Timer
}
