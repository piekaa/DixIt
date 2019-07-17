package game

import (
	"dixit/game/state"
	"errors"
)

func (this *game) Pull(roomName string) (*PullResponse, error) {

	if !this.roomRepository.Has(roomName) {
		return nil, errors.New("room does not exist")
	}
	r := this.roomRepository.Get(roomName)
	var payload interface{}
	switch r.GameState {
	case state.LOBBY:
		payload = &LobbyPayload{r.Players}
	case state.WHO_FIRST:
		payload = &LobbyPayload{r.Players}
	case state.CHOOSE_CARDS:
		payload = &GameplayPayload{r.Players, r.ActivePlayer}
	case state.ROUND_RESULT:
		payload = &GameplayPayload{r.Players, r.ActivePlayer}
	}
	return &PullResponse{r.GameState, payload}, nil
}
