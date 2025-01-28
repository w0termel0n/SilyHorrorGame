package silygame

import (
	//"math"
	//"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/w0termel0n/SilyHorrorGame/assets"
)

type Player struct {
	game     *Game
	position Vector
	sprite   *ebiten.Image
}

// creates player object which has an image and positon that can changed to move it
func NewPlayer(game *Game, img string) *Player {
	var sprite *ebiten.Image
	if img == "sprite" {
		sprite = assets.PlayerSprite
	} else {
		sprite = assets.EndSprite
	}

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: screenWidth/2 - halfW,
		Y: screenHeight/2 - halfH,
	}

	return &Player{
		game:     game,
		position: pos,
		sprite:   sprite,
	}
}

// adds sprite to the screen at the center
func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}

// adds a hitbox with collision detection around the sprite
func (p *Player) Collider() Rect {
	bounds := p.sprite.Bounds()

	return NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
