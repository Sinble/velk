package structs

import "velk/src/interfaces"

type Mob struct {
	Id int
	Name string
	Room *Room
	State string
	Targets []*Player
	Items		map[string]interfaces.ItemInterface

	Health int
	MaxHealth int
}

func (mob Mob) SendToPlayer(message string) {

}

func (mob Mob) ReadFromPlayer() (string, error) {

	return "", nil
}

func(mob Mob) SendPlayerPrompt() {

}