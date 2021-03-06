package game

import (
	"dixit/game/player"
	"dixit/game/room"
	"dixit/game/state"
)

func (this *game) Start(roomName string) *room.Room {

	newRoom := &room.Room{RoomState: room.NEW, Name: roomName, Players: map[string]*player.Player{}, PlayersByPosition: map[int]*player.Player{}, ActivePlayer: ""}
	newRoom.GameState = state.LOBBY

	this.roomRepository.Add(newRoom)
	return newRoom
}
