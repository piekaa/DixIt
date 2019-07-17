package game

import (
	"errors"
)

func (this *game) Ready(roomName, playerName string) error {

	if !this.roomRepository.Has(roomName) {
		return errors.New("Room does not exist")
	}
	r := this.roomRepository.Get(roomName)
	p, has := r.Players[playerName]
	if !has {
		return errors.New("Player does not exist in room")
	}

	p.ReadyToStart = true

	this.startIfAllReady(r)

	return nil
}
