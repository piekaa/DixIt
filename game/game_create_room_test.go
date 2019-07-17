package game

import (
	"dixit/game/room"
	"testing"
)

func TestGame_Start(t *testing.T) {
	g := NewGame()

	r := g.Start("room1")
	if r.RoomState != room.NEW {
		t.Error("RoomState should be NEW")
	}
	if r.Name != "room1" {
		t.Error("Romm name should be room1")
	}

	r = g.Start("room1")
	if r.RoomState != room.NEW {
		t.Error("RoomState should be NEW")
	}
	if r.Name != "room1" {
		t.Error("Romm name should be room1")
	}

	r = g.Start("room2")
	if r.RoomState != room.NEW {
		t.Error("RoomState should be NEW")
	}
	if r.Name != "room2" {
		t.Error("Romm name should be room2")
	}
}
