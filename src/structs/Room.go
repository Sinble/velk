package structs

import (
	"sync"
	"velk/src/interfaces"
)

type Room struct {
	ID          int
	Name        string
	Description string
	players     map[int]*Player
	mobs		map[int]*Mob
	Items		map[string]interfaces.ItemInterface
	NorthExitID string
	EastExitID  string
	SouthExitID string
	WestExitID  string
	NorthExit   *Room
	EastExit    *Room
	SouthExit   *Room
	WestExit    *Room
	playerMutex sync.Mutex
	mobMutex sync.Mutex
}

func (r Room) New(id int) *Room {

	room := &Room{
		ID:          id,
		Name:        "A room",
		Description: "A plain looking room",
		players:     make(map[int]*Player),
		mobs:		 make(map[int]*Mob),
		Items:       make(map[string]interfaces.ItemInterface),
		NorthExitID: "",
		EastExitID:  "",
		SouthExitID: "",
		WestExitID:  "",
		NorthExit:   nil,
		EastExit:    nil,
		SouthExit:   nil,
		WestExit:    nil,
		playerMutex: sync.Mutex{},
		mobMutex: sync.Mutex{},
	}

	return room
}

func (r *Room) AddPlayer(player *Player) {
	r.playerMutex.Lock()
	r.players[player.Id] = player
	r.playerMutex.Unlock()
}

func (r *Room) RemovePlayer(player *Player) {
	r.playerMutex.Lock()
	delete(r.players, player.Id)
	r.playerMutex.Unlock()
}

func (r *Room) AddMob(mob *Mob) {
	r.mobMutex.Lock()
	r.mobs[mob.Id] = mob
	r.mobMutex.Unlock()
}

func (r *Room) RemoveMob(mob *Mob) {
	r.mobMutex.Lock()
	delete(r.mobs, mob.Id)
	r.mobMutex.Unlock()
}

func(r *Room) MobCheck(mobId int) (*Mob, *sync.Mutex) {
	r.mobMutex.Lock()
	mob := r.mobs[mobId]

	return mob, &r.mobMutex
}

func(r *Room) GetPlayers() map[int]*Player {
	return r.players
}

func(r *Room) GetMobs() map[int]*Mob {
	return r.mobs
}