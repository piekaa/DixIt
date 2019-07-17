package game

import "dixit/game/player"

type PullResponse struct {
	GameState string
	Payload   interface{}
}

type LobbyPayload struct {
	Players map[string]*player.Player
}

type GameplayPayload struct {
	Players      map[string]*player.Player
	ActivePlayer string
}
