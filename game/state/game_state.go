package state

//game states
const (
	LOBBY        = "lobby"
	WHO_FIRST    = "whoFirst"
	CHOOSE_CARDS = "chooseCards"
	//in case 2 or more players claim same card is theirs
	ROUND_FALSE_RESULT = "roundFalseResult"
	ROUND_RESULT       = "roundResult"
)
