package game

import (
	"dixit/game/player"
	"dixit/game/room"
	"dixit/game/state"
)

type Game interface {
	Start(roomName string) *room.Room
	//false if playerName is taken
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

	if r.GameState == state.LOBBY {
		for _, p := range r.Players {
			r.PlayersByPosition[p.Position] = p
		}
	}

	for _, p := range r.Players {
		p.ReadyToStart = false
		p.MyCard = 0
		p.Vote = 0
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
	activePlayerCardVoteCount := 0

	cardPositionPlayerMap := map[int]*player.Player{}
	myCards := map[int]bool{}

	for _, p := range r.Players {

		if _, ok := myCards[p.MyCard]; ok {
			r.GameState = state.ROUND_FALSE_RESULT
			return
		}

		myCards[p.MyCard] = true

		cardPositionPlayerMap[p.MyCard] = p
		if p.Vote == 0 {
			return
		}
		if p.Name == r.ActivePlayer {
			continue
		}
		if p.Vote == activePlayerCard {
			activePlayerCardVoteCount++
		}
	}

	r.GameState = state.ROUND_RESULT

	if activePlayerCardVoteCount == 0 || activePlayerCardVoteCount == len(r.Players)-1 {
		for _, p := range r.Players {
			if p.Name == r.ActivePlayer {
				continue
			}
			p.Score += 2

			if p.Vote != activePlayerCard {
				cardPositionPlayerMap[p.Vote].Score++
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
				cardPositionPlayerMap[p.Vote].Score++
			}
		}
	}
}
