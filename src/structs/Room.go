package structs

type Room struct {
	Id int
	Name string
	Description string
	Players []Player
	NorthExit int
	EastExit int
	SouthExit int
	WestExit int
}

