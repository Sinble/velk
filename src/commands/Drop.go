package commands

import (
	"velk/src/interfaces"
	"velk/src/structs"
)

type Drop struct {

}

func (a *Drop) Action(playerInterface interfaces.PlayerInterface, command string, commandOptions ...string) {
	player := playerInterface.(*structs.Player)

	for id, item := range player.Items {
		if item.GetName() == commandOptions[0] {
			delete(player.Items, id)
			player.Room.Items[id] = item
			return
		}
	}

	player.SendToPlayer("you don't see that anywhere/r/n")
}
