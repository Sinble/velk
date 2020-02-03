package structs

type Room struct {
	Id int
	Name string
	Description string
	Players []*Player
	NorthExitId string
	EastExitId string
	SouthExitId string
	WestExitId string
	NorthExit *Room
	EastExit *Room
	SouthExit *Room
	WestExit *Room
}

