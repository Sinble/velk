package structs

import (
	"velk/src/interfaces"
	"github.com/lithammer/shortuuid/v3"
)

type Container struct {
	ID string
	IID int
	Name string
	Items map[string]interfaces.ItemInterface
}

func (i Container) New() *Container {
	container := &Container{
		ID: shortuuid.New(),
		Items: make(map[string]interfaces.ItemInterface),
	}

	return container
}

func (i *Container) GetName() string {
	return i.Name
}