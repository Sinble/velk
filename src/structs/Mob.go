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
