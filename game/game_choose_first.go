package game

import (
	"dixit/game/state"
	"errors"
)

func (this *game) ChooseFirst(roomName, playerName string) (bool, error) {
	if !this.roomRepository.Has(roomName) {
		return false, errors.New("Room does not exist")
	}
	r := this.roomRepository.Get(roomName)

	_, has := r.Players[playerName]
	if !has {
		return false, errors.New("Player does not exist in room")
	}

	if r.ActivePlayer != "" {
		return false, nil
	}

	r.ActivePlayer = playerName
	r.GameState = state.CHOOSE_CARDS

	return true, nil
}
