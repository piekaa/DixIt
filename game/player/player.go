package player

type Player struct {
	Name         string
	ReadyToStart bool
	Position     int
	DidChoose    bool
	Vote         int
	MyCard       int
	Score        int
}
