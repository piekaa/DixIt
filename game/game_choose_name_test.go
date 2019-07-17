package game

import "testing"

func TestGame_ChooseName(t *testing.T) {
	g := NewGame()

	g.Start("r1")
	g.Start("r2")

	_, err := g.ChooseName("r3", "n1")

	if err == nil {
		t.Error("error is nil")
	}

	suc, err := g.ChooseName("r1", "n1")

	if err != nil {
		t.Error(err)
	}

	if !suc {
		t.Fail()
	}

	suc, err = g.ChooseName("r1", "n2")

	if err != nil {
		t.Error(err)
	}

	if !suc {
		t.Fail()
	}

	suc, err = g.ChooseName("r1", "n1")

	if err != nil {
		t.Error(err)
	}

	if suc {
		t.Fail()
	}

	suc, err = g.ChooseName("r2", "n1")

	if err != nil {
		t.Error(err)
	}

	if !suc {
		t.Fail()
	}
}
