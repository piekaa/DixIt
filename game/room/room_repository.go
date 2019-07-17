package room

type RoomRepository interface {
	Add(room *Room)
	Has(name string) bool
	Get(name string) *Room
}

type roomRepository struct {
	rooms map[string]*Room
}

func (this *roomRepository) Add(room *Room) {
	this.rooms[room.Name] = room
}

func (this *roomRepository) Has(name string) bool {
	_, has := this.rooms[name]
	return has
}

func (this *roomRepository) Get(name string) *Room {
	return this.rooms[name]
}

func NewRoomRepository() RoomRepository {
	return &roomRepository{map[string]*Room{}}
}
