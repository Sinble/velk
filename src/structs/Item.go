package structs
import "github.com/lithammer/shortuuid/v3"

type Item struct {
	ID string
	IID int
	Name string
}

func (i Item) New() *Item {
	
	item := &Item{
		ID: shortuuid.New(),
	}

	return item
}

func (i *Item) GetName() string {
	return i.Name
}
