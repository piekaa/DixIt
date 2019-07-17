package game

import (
	"dixit/game/state"
	"errors"
	"strconv"
)

func (this *game) ChooseCards(roomName, playerName string, myCard, myType int) error {

	if !this.roomRepository.Has(roomName) {
		return errors.New("room does not exist")
	}
	r := this.roomRepository.Get(roomName)

	if r.GameState != state.CHOOSE_CARDS {
		return errors.New("choosing card in " + r.GameState + " state is not allowed")
	}

	p, has := r.Players[playerName]
	if !has {
		return errors.New("player does not exist in room")
	}

	if myCard <= 0 || myCard > len(r.Players) {
		return errors.New("player choosed incorrect myCard value=" + strconv.Itoa(myCard))
	}

	if myType <= 0 || myType > len(r.Players) {
		return errors.New("player choosed incorrect myType value=" + strconv.Itoa(myCard))
	}

	if r.ActivePlayer != p.Name && myCard == myType {
		return errors.New("player can't vote for his card")
	}

	p.Vote = myType
	p.MyCard = myCard
	p.DidChoose = true

	this.calculateScoreIfAllPlayersVoted(r)

	return nil
}
