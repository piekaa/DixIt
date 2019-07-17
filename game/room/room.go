package room

import "dixit/game/player"

const (
	NEW     = iota
	PLAYING = iota
)

type Room struct {
	RoomState         int
	GameState         string
	Name              string
	Players           map[string]*player.Player
	PlayersByPosition map[int]*player.Player
	ActivePlayer      string
}
