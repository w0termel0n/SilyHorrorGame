package silygame

import (
	//"math"
	//"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Room struct {
	game           *Game
	name           string
	is_hall        bool
	background     *ebiten.Image
	adj_rooms      []*Room
	adj_directions []string
}

// constructs room struct
func NewRoom(game *Game, name string, is_hall bool, background *ebiten.Image) *Room {
	return &Room{
		game:       game,
		name:       name,
		is_hall:    is_hall,
		background: background,
	}
}

func LinkRooms(room *Room, adj *Room, dir string) {
	room.adj_rooms = append(room.adj_rooms, adj)
	room.adj_directions = append(room.adj_directions, dir)
}
