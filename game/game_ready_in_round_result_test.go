package game

import (
	"dixit/game/state"
	"testing"
)

func TestGame_READY_IN_STATE_RESULT(t *testing.T) {

	g := prepareGameWith3Players(t)
	err := g.ChooseCards("r1", "n1", 1, 1)
	if err != nil {
		t.Error(err)
	}
	err = g.ChooseCards("r1", "n2", 2, 1)
	if err != nil {
		t.Error(err)
	}
	err = g.ChooseCards("r1", "n3", 3, 1)
	if err != nil {
		t.Error(err)
	}
	p, err := g.Pull("r1")

	if err != nil {
		t.Error(err)
	}

	if p.GameState != state.ROUND_RESULT {
		t.Error()
	}

	err = g.Ready("r1", "n1")

	if err != nil {
		t.Error(err)
	}

	err = g.Ready("r1", "n2")

	if err != nil {
		t.Error(err)
	}
	err = g.Ready("r1", "n3")

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
}
