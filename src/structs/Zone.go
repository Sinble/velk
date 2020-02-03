package structs

type Zone struct {
	Id   int
	Name string
	Creator string
	Rooms map[int]*Room
}