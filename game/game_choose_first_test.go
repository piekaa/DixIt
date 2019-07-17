package game

import "testing"

func TestGame_ChooseFirst(t *testing.T) {

	g := NewGame()
	g.Start("r1")
	g.ChooseName("r1", "n1")
	g.ChooseName("r1", "n2")
	g.ChooseName("r1", "n3")
	g.Ready("r1", "n1")
	g.Ready("r1", "n2")
	g.Ready("r1", "n3")

	_, err := g.ChooseFirst("r2", "n1")

	if err == nil {
		t.Error()
	}

	_, err = g.ChooseFirst("r1", "n4")

	if err == nil {
		t.Error()
	}

	first, err := g.ChooseFirst("r1", "n1")

	if err != nil {
		t.Error(err)
	}

	if !first {
		t.Error()
	}

	first, err = g.ChooseFirst("r1", "n3")

	if err != nil {
		t.Error(err)
	}

	if first {
		t.Error()
	}

}
