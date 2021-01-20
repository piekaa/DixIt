package player

type Player struct {
	Name         string `json:"playerName"`
	ReadyToStart bool   `json:"readyToStart"`
	Position     int    `json:"position"`
	DidChoose    bool   `json:"didChoose"`
	Vote         int    `json:"vote"`
	MyCard       int    `json:"myCard"`
	Score        int    `json:"score"`
}
