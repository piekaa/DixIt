package game

import (
	"dixit/game/player"
	"fmt"
	"testing"
)

func TestT(t *testing.T) {

	w := &wrap{&player.Player{}}

	s, ok := w.lala.(*player.Player)

	fmt.Println(ok)

	fmt.Println(s)

}

func TestMap(t *testing.T) {

	m := map[string]*player.Player{}

	m["test"] = &player.Player{}

	p := m["test"]

	p.ReadyToStart = true

	fmt.Println(m["test"].ReadyToStart)

}

type wrap struct {
	lala interface{}
}
