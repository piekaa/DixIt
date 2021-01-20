package game

import (
	"dixit/game/state"
	"testing"
)

//todo once it's pass other time it fails
func TestGame_Pull_IN_FALSE_STATE_RESULT(t *testing.T) {

	g := prepareGameWith3Players(t)

	err := g.ChooseCards("r1", "n1", 1, 1)
	if err != nil {
		t.Error(err)
	}
	err = g.ChooseCards("r1", "n2", 1, 2)
	if err != nil {
		t.Error(err)
	}
	err = g.ChooseCards("r1", "n3", 2, 1)
	if err != nil {
		t.Error(err)
	}
	p, err := g.Pull("r1")

	if err != nil {
		t.Error()
	}

	if p == nil {
		t.Error()
	}

	if p.GameState != state.ROUND_FALSE_RESULT {
		t.Error()
	}
}
