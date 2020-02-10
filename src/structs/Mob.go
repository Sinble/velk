package structs

type Mob struct {
	Id int
	Name string
	Room *Room
	State string
	Targets []*Player

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