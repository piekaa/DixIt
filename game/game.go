package game

import (
	"dixit/game/player"
	"dixit/game/room"
	"dixit/game/state"
)

type Game interface {
	Start(roomName string) *room.Room
	//false if name is taken
	ChooseName(roomName, playerName string) (success bool, err error)
	Ready(roomName, playerName string) error
	ChooseFirst(roomName, playerName string) (bool, error)
	ChooseCards(roomName, playerName string, myCard, myType int) error

	Pull(roomName string) (*PullResponse, error)
	calculateScoreIfAllPlayersVoted(room *room.Room)
}

type game struct {
	roomRepository room.RoomRepository
}

func NewGame() Game {
	return &game{room.NewRoomRepository()}
}

func (this *game) areAllReady(r *room.Room) bool {
	for _, p := range r.Players {

		if !p.ReadyToStart {
			return false
		}
	}
	return true
}

func (this *game) startIfAllReady(r *room.Room) {

	if !this.areAllReady(r) {
		return
	}

	for _, p := range r.Players {
		p.ReadyToStart = false
	}

	nextState := map[string]string{}
	nextState[state.LOBBY] = state.WHO_FIRST
	nextState[state.ROUND_RESULT] = state.CHOOSE_CARDS
	nextState[state.ROUND_FALSE_RESULT] = state.CHOOSE_CARDS

	if r.GameState == state.ROUND_RESULT {
		currentActivePlayerPosition := r.Players[r.ActivePlayer].Position
		nextActivePlayerPosition := 0
		if currentActivePlayerPosition == len(r.Players) {
			nextActivePlayerPosition = 1
		} else {
			nextActivePlayerPosition = currentActivePlayerPosition + 1
		}
		r.ActivePlayer = r.PlayersByPosition[nextActivePlayerPosition].Name
	}

	r.GameState = nextState[r.GameState]

}

func (this *game) calculateScoreIfAllPlayersVoted(r *room.Room) {

	activePlayerCard := r.Players[r.ActivePlayer].MyCard
	activePlayerCardVoteCound := 0

	positionPlayerMap := map[int]*player.Player{}
	myCards := map[int]bool{}

	for _, p := range r.Players {

		if _, ok := myCards[p.MyCard]; ok {
			r.GameState = state.ROUND_FALSE_RESULT
			return
		}

		myCards[p.MyCard] = true

		positionPlayerMap[p.Position] = p
		if p.Vote == 0 {
			return
		}
		if p.Name == r.ActivePlayer {
			continue
		}
		if p.Vote == activePlayerCard {
			activePlayerCardVoteCound++
		}
	}

	r.GameState = state.ROUND_RESULT

	if activePlayerCardVoteCound == 0 || activePlayerCardVoteCound == len(r.Players)-1 {
		for _, p := range r.Players {
			if p.Name == r.ActivePlayer {
				continue
			}
			p.Score += 2

			if p.Vote != activePlayerCard {
				positionPlayerMap[p.Vote].Score++
			}
		}
	} else {
		for _, p := range r.Players {
			if p.Name == r.ActivePlayer {
				continue
			}

			if p.Vote == activePlayerCard {
				p.Score += 3
				r.Players[r.ActivePlayer].Score += 3
			} else {
				positionPlayerMap[p.Vote].Score++
			}
		}
	}
}
