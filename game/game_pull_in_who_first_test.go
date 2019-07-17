package game

import (
	"dixit/game/state"
	"testing"
)

func TestGame_Pull_IN_WHO_FIRST(t *testing.T) {

	g := NewGame()
	g.Start("r1")
	g.ChooseName("r1", "n1")
	g.ChooseName("r1", "n2")
	g.ChooseName("r1", "n3")
	g.Ready("r1", "n1")
	g.Ready("r1", "n2")
	g.Ready("r1", "n3")

	p, err := g.Pull("r1")

	if err != nil {
		t.Error()
	}

	_, ok := p.Payload.(*LobbyPayload)

	if !ok {
		t.Error()
	}

	if p.GameState != state.WHO_FIRST {
		t.Error()
	}

	g.ChooseFirst("r1", "n3")

	p, err = g.Pull("r1")

	if err != nil {
		t.Error(err)
	}

	payload, ok := p.Payload.(*GameplayPayload)

	if !ok {
		t.Error()
	}

	if p.GameState != state.CHOOSE_CARDS {
		t.Error()
	}

	if payload.ActivePlayer != "n3" {
		t.Error()
	}

}
