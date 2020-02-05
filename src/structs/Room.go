package structs

import "sync"

type Room struct {
	ID          int
	Name        string
	Description string
	players     map[int]*Player
	NorthExitID string
	EastExitID  string
	SouthExitID string
	WestExitID  string
	NorthExit   *Room
	EastExit    *Room
	SouthExit   *Room
	WestExit    *Room
	playerMutex sync.Mutex
}

func (r Room) New(id int) *Room {

	room := &Room{
		ID:          id,
		Name:        "A room",
		Description: "A plain looking room",
		players:     make(map[int]*Player, 0),
		NorthExitID: "",
		EastExitID:  "",
		SouthExitID: "",
		WestExitID:  "",
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