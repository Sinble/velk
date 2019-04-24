package interfaces

import "velk/src/structs"

type CommandInterface interface {
	Action(structs.Player, string, ...string)
}
