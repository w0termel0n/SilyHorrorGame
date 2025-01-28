package silygame

import (
	//"fmt"
	//"image/color"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/w0termel0n/SilyHorrorGame/assets"
)

const (
	screenWidth  = 800
	screenHeight = 600

	baseVelocity = 2.5
	duration     = 15 * time.Second //duration for initializing a timer
)

type Game struct {
	player   *Player
	xVel     float64
	yVel     float64
	velTimer *Timer
	walls    []Rect //i made the walls of the game window hitboxes to detect them
	gameOver bool
}

func NewGame() *Game {
	g := &Game{
		xVel:     baseVelocity,
		yVel:     baseVelocity,
		velTimer: NewTimer(duration),
		gameOver: false,
	}

	g.player = NewPlayer(g, "sprite")

	g.walls = append(g.walls, NewRect(0, 0, screenWidth, 1))
	g.walls = append(g.walls, NewRect(screenWidth-1, 0, 1, screenHeight))
	g.walls = append(g.walls, NewRect(0, screenHeight-1, screenWidth, 1))
	g.walls = append(g.walls, NewRect(0, 0, 1, screenHeight))

	return g
}

// the main game loop
// the RunGame function in main runs this every tick
func (g *Game) Update() error {
	g.velTimer.Update()
	//when the timer runs out, run this
	if g.velTimer.IsReady() && !g.gameOver {
		g.player.sprite.Deallocate()
		g.player = NewPlayer(g, "end")
		g.velTimer.Reset()
		//when it runs out the second time
	} else if g.velTimer.IsReady() && g.gameOver {
		os.Exit(0)
	}

	//makes the lil guy bounce when hitting a wall
	for i, w := range g.walls {
		if i%2 == 1 && g.player.Collider().Intersects(w) {
			g.xVel = -(g.xVel)
		}
		if i%2 == 0 && g.player.Collider().Intersects(w) {
			g.yVel = -(g.yVel)
		}
	}

	g.player.position.X += g.xVel
	g.player.position.Y += g.yVel

	return nil
}

// the RunGame func also runs this every tick
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(assets.Background, nil)
	g.player.Draw(screen)
}

// RunGame runs this every so often, like when creating the game window
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
