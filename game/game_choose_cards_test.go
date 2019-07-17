package game

import (
	"dixit/game/state"
	"testing"
)

func TestGame_NonActivePlayerCantChooseHisCard(t *testing.T) {

	g := prepareGameWith3Players(t)

	err := g.ChooseCards("r1", "n1", 1, 1)

	if err != nil {
		t.Error()
	}

	err = g.ChooseCards("r1", "n2", 1, 1)

	if err == nil {
		t.Error()
	}
}

func TestGame_PlayerVoteForIncorrectCard(t *testing.T) {

	cases := []struct {
		myCard, myType int
	}{
		{1, 5},
		{1, -1},
		{0, 2},
		{-1, 2},
		{-1, -2},
		{4, 2}}

	for _, v := range cases {
		g := prepareGameWith3Players(t)
		err := g.ChooseCards("r1", "n2", v.myCard, v.myType)
		if err == nil {
			t.Error("Player choosed incorrect myCard=", v.myCard, " or myType=", v.myType)
		}
	}

}

func TestGame_TwoPeopleVoteForSameMyCard(t *testing.T) {

	g := prepareGameWith3Players(t)

	err := g.ChooseCards("r1", "n1", 1, 1)
	if err != nil {
		t.Error(err)
	}

	err = g.ChooseCards("r1", "n2", 2, 1)
	if err != nil {
		t.Error(err)
	}

	err = g.ChooseCards("r1", "n3", 2, 1)
	if err != nil {
		t.Error(err)
	}

	p, err := g.Pull("r1")

	if err != nil {
		t.Error(err)
	}

	if p.GameState != state.ROUND_FALSE_RESULT {
		t.Error("game should be in round false result state")
	}

}

func TestGame_AllVotedCorrectly(t *testing.T) {

	g := prepareGameWith3Players(t)

	err := g.ChooseCards("r1", "n1", 1, 1)
	if err != nil {
		t.Error(err)
	}

	err = g.ChooseCards("r1", "n2", 3, 2)
	if err != nil {
		t.Error(err)
	}

	err = g.ChooseCards("r1", "n3", 2, 1)
	if err != nil {
		t.Error(err)
	}

	p, err := g.Pull("r1")

	if err != nil {
		t.Error(err)
	}

	if p.GameState != state.ROUND_RESULT {
		t.Error("game should be in round result state")
	}
}

func TestGame_CanChangeVoteIfNotEveryoneVoted(t *testing.T) {
	g := prepareGameWith3Players(t)

	err := g.ChooseCards("r1", "n1", 1, 1)
	if err != nil {
		t.Error(err)
	}

	err = g.ChooseCards("r1", "n2", 3, 2)
	if err != nil {
		t.Error(err)
	}

	err = g.ChooseCards("r1", "n2", 3, 1)
	if err != nil {
		t.Error(err)
	}

	err = g.ChooseCards("r1", "n3", 2, 1)
	if err != nil {
		t.Error(err)
	}

	p, err := g.Pull("r1")

	if err != nil {
		t.Error(err)
	}

	if p.GameState != state.ROUND_RESULT {
		t.Error("game should be in round result state")
	}
}

func TestGame_CantChangeVoteIfEveryoneVoted(t *testing.T) {
	g := prepareGameWith3Players(t)

	err := g.ChooseCards("r1", "n1", 1, 1)
	if err != nil {
		t.Error(err)
	}

	err = g.ChooseCards("r1", "n2", 3, 2)
	if err != nil {
		t.Error(err)
	}

	err = g.ChooseCards("r1", "n3", 2, 1)
	if err != nil {
		t.Error(err)
	}

	err = g.ChooseCards("r1", "n2", 3, 1)
	if err == nil {
		t.Error("choosing cards after all did it should be not allowed")
	}
}
