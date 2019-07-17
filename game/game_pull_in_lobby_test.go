package game

import (
	"dixit/game/state"
	"testing"
)

func TestGame_Pull_IN_LOBBY_ON_CHOOSE_NAME(t *testing.T) {

	g := NewGame()
	g.Start("r1")

	_, err := g.Pull("r2")

	if err == nil {
		t.Error()
	}

	p, err := g.Pull("r1")

	if err != nil {
		t.Error()
	}

	if p.GameState != state.LOBBY {
		t.Error()
	}
	payload, ok := p.Payload.(*LobbyPayload)

	if !ok {
		t.Error()
	}

	if len(payload.Players) != 0 {
		t.Error()
	}

	g.ChooseName("r1", "n1")
	g.ChooseName("r1", "n2")
	g.ChooseName("r1", "n3")

	p, err = g.Pull("r1")

	if err != nil {
		t.Error()
	}

	if p.GameState != state.LOBBY {
		t.Error()
	}
	payload, ok = p.Payload.(*LobbyPayload)

	if !ok {
		t.Error()
	}

	if len(payload.Players) != 3 {
		t.Error()
	}

	if payload.Players["n1"].ReadyToStart || payload.Players["n1"].Position != 1 {
		t.Error()
	}

	if payload.Players["n2"].ReadyToStart || payload.Players["n2"].Position != 2 {
		t.Error()
	}

	if payload.Players["n3"].ReadyToStart || payload.Players["n3"].Position != 3 {
		t.Error()
	}
}

func TestGame_Pull_IN_LOBBY_ON_READY(t *testing.T) {

	g := NewGame()
	g.Start("r1")

	g.ChooseName("r1", "n1")
	g.ChooseName("r1", "n2")
	g.ChooseName("r1", "n3")

	err := g.Ready("r1", "n1")

	if err != nil {
		t.Error(err)
	}

	err = g.Ready("r1", "n3")

	if err != nil {
		t.Error(err)
	}

	p, err := g.Pull("r1")

	if err != nil {
		t.Error()
	}

	if p.GameState != state.LOBBY {
		t.Error()
	}
	payload, ok := p.Payload.(*LobbyPayload)

	if !ok {
		t.Error()
	}

	if len(payload.Players) != 3 {
		t.Error()
	}

	if !payload.Players["n1"].ReadyToStart {
		t.Error()
	}

	if payload.Players["n2"].ReadyToStart {
		t.Error()
	}

	if !payload.Players["n3"].ReadyToStart {
		t.Error()
	}

	g.Ready("r1", "n2")

	p, err = g.Pull("r1")

	if err != nil {
		t.Error()
	}

	if p.GameState != state.WHO_FIRST {
		t.Error()
	}
}
