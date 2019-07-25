package game

import "dixit/game/player"

type PullResponse struct {
	GameState string      `json:"gameState"`
	Payload   interface{} `json:"payload"`
}

type LobbyPayload struct {
	Players map[string]*player.Player `json:"players"`
}

type GameplayPayload struct {
	Players      map[string]*player.Player `json:"players"`
	ActivePlayer string                    `json:"activePlayer"`
}
