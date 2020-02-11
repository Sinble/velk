package commands

import (
	"velk/src/interfaces"
	"velk/src/structs"
)

type Get struct {

}

func (a *Get) Action(playerInterface interfaces.PlayerInterface, command string, commandOptions ...string) {
	player := playerInterface.(*structs.Player)

	for id, item := range player.Room.Items {
		if item.GetName() == commandOptions[0] {
			delete(player.Room.Items, id)
			player.Items[id] = item
			return
		}
	}

	player.SendToPlayer("you don't see that anywhere/r/n")
}
