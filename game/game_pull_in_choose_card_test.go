package game

import (
	"dixit/game/state"
	"testing"
)

func TestGame_PullInChooseCard(t *testing.T) {

	g := NewGame()
	g.Start("r1")
	g.ChooseName("r1", "n1")
	g.ChooseName("r1", "n2")
	g.ChooseName("r1", "n3")
	g.Ready("r1", "n1")
	g.Ready("r1", "n2")
	g.Ready("r1", "n3")

	g.ChooseFirst("r1", "n1")

	p, err := g.Pull("r1")

	if err != nil {
		t.Error(err)
	}

	if p.GameState != state.CHOOSE_CARDS {
		t.Error()
	}

	payload, ok := p.Payload.(*GameplayPayload)

	if !ok {
		t.Error()
	}

	if payload.ActivePlayer != "n1" {
		t.Error()
	}

	if payload.Players["n1"].DidChoose {
		t.Error()
	}

	if payload.Players["n2"].DidChoose {
		t.Error()
	}

	if payload.Players["n3"].DidChoose {
		t.Error()
	}

	g.ChooseCards("r1", "n1", 1, 1)

	p, err = g.Pull("r1")

	if err != nil {
		t.Error(err)
	}

	if p.GameState != state.CHOOSE_CARDS {
		t.Error()
	}

	payload, ok = p.Payload.(*GameplayPayload)

	if !ok {
		t.Error()
	}

	if payload.ActivePlayer != "n1" {
		t.Error()
	}

	if !payload.Players["n1"].DidChoose {
		t.Error()
	}

	if payload.Players["n2"].DidChoose {
		t.Error()
	}

	if payload.Players["n3"].DidChoose {
		t.Error()
	}

	err = g.ChooseCards("r1", "n3", 3, 2)

	if err != nil {
		t.Error(err)
	}

	p, err = g.Pull("r1")

	if err != nil {
		t.Error(err)
	}

	if p.GameState != state.CHOOSE_CARDS {
		t.Error()
	}

	payload, ok = p.Payload.(*GameplayPayload)

	if !ok {
		t.Error()
	}

	if payload.ActivePlayer != "n1" {
		t.Error()
	}

	if !payload.Players["n1"].DidChoose {
		t.Error()
	}

	if payload.Players["n2"].DidChoose {
		t.Error()
	}

	if !payload.Players["n3"].DidChoose {
		t.Error()
	}

	g.ChooseCards("r1", "n2", 2, 3)

	payload, ok = p.Payload.(*GameplayPayload)

	if !ok {
		t.Error()
	}

	if payload.ActivePlayer != "n1" {
		t.Error()
	}

	if !payload.Players["n1"].DidChoose {
		t.Error()
	}

	if !payload.Players["n2"].DidChoose {
		t.Error()
	}

	if !payload.Players["n3"].DidChoose {
		t.Error()
	}

	p, err = g.Pull("r1")

	if err != nil {
		t.Error(err)
	}

	if p.GameState != state.ROUND_RESULT {
		t.Error()
	}
}
