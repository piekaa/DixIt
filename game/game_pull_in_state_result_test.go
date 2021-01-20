package game

import (
	"dixit/game/state"
	"testing"
)

func TestGame_Pull_IN_STATE_RESULT(t *testing.T) {

	cases := []struct {
		p1myCard, p1myType, p2myCard, p2myType, p3myCard, p3myType, p1score, p2score, p3score int
	}{
		{1, 1, 2, 3, 3, 2, 0, 3, 3}, // all miss
		{2, 2, 1, 3, 3, 1, 0, 3, 3}, // all miss
		{3, 3, 1, 2, 2, 1, 0, 3, 3}, // all miss
		{1, 1, 2, 1, 3, 1, 0, 2, 2}, // all hit
		{1, 3, 2, 1, 3, 2, 3, 4, 0},
		{1, 1, 2, 1, 3, 2, 3, 4, 0},
		{1, 1, 2, 3, 3, 1, 3, 0, 4}} // one hit
	for i, v := range cases {
		g := prepareGameWith3Players(t)
		err := g.ChooseCards("r1", "n1", v.p1myCard, v.p1myType)
		if err != nil {
			t.Error(err)
		}
		err = g.ChooseCards("r1", "n2", v.p2myCard, v.p2myType)
		if err != nil {
			t.Error(err)
		}
		err = g.ChooseCards("r1", "n3", v.p3myCard, v.p3myType)
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

		payload, ok := p.Payload.(*GameplayPayload)

		if !ok {
			t.Error()
		}

		if payload.Players["n1"].Score != v.p1score {
			t.Error("iteration: ", i, " n1 score should be ", v.p1score, "but is", payload.Players["n1"].Score)
		}

		if payload.Players["n2"].Score != v.p2score {
			t.Error("iteration: ", i, " n2 score should be ", v.p2score, "but is", payload.Players["n2"].Score)
		}

		if payload.Players["n3"].Score != v.p3score {
			t.Error("iteration: ", i, " n3 score should be ", v.p3score, "but is", payload.Players["n3"].Score)
		}
	}
}
