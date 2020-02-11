package structs

import "velk/src/interfaces"

type Container struct {
	ID int
	Name string
	Items map[int]interfaces.ItemInterface
}

func (i *Container) GetName() string {
	return i.Name
}