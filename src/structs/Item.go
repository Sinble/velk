package structs

type Item struct {
	ID int
	Name string
}

func (i *Item) GetName() string {
	return i.Name
}
