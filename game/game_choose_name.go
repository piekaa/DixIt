package game

import (
	"dixit/game/player"
	"errors"
)

func (this *game) ChooseName(roomName, playerName string) (bool, error) {

	if !this.roomRepository.Has(roomName) {
		return false, errors.New("Room does not exist")
	}
	r := this.roomRepository.Get(roomName)
	_, has := r.Players[playerName]
	if has {
		return false, nil
	}
	r.Players[playerName] = &player.Player{Name: playerName, Position: len(r.Players) + 1}
	return true, nil
}
