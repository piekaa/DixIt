package game

import "testing"

func prepareGameWith3Players(t *testing.T) Game {
	g := NewGame()
	g.Start("r1")
	g.ChooseName("r1", "n1")
	g.ChooseName("r1", "n2")
	g.ChooseName("r1", "n3")
	g.Ready("r1", "n1")
	g.Ready("r1", "n2")
	g.Ready("r1", "n3")

	g.ChooseFirst("r1", "n1")

	err := g.ChooseCards("r2", "n1", 2, 1)

	if err == nil {
		t.Error()
	}

	err = g.ChooseCards("r1", "n5", 2, 1)

	if err == nil {
		t.Error()
	}

	return g
}
