package structs

import "sync"

type Room struct {
	Id          int
	Name        string
	Description string
	players     map[int]*Player
	NorthExitId string
	EastExitId  string
	SouthExitId string
	WestExitId  string
	NorthExit   *Room
	EastExit    *Room
	SouthExit   *Room
	WestExit    *Room
	playerMutex sync.Mutex
}

func (r Room) New(id int) Room {

	room := Room{
		Id:          id,
		Name:        "A room",
		Description: "A plain looking room",
		players:     make(map[int]*Player, 0),
		NorthExitId: "",
		EastExitId:  "",
		SouthExitId: "",
		WestExitId:  "",
		NorthExit:   nil,
		EastExit:    nil,
		SouthExit:   nil,
		WestExit:    nil,
		playerMutex: sync.Mutex{},
	}

	return room
}

func (r Room) AddPlayer(player *Player) {
	r.playerMutex.Lock()
	r.players[player.Id] = player
	r.playerMutex.Unlock()
}

func (r Room) RemovePlayer(player *Player) {
	r.playerMutex.Lock()
	delete(r.players, player.Id)
	r.playerMutex.Unlock()
}

func(r Room) GetPlayers() map[int]*Player {
	return r.players
}