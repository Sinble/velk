package commands

import (
	"fmt"
	"velk/src/interfaces"
	"velk/src/structs"
)

type Inventory struct {

}

func (a *Inventory) Action(playerInterface interfaces.PlayerInterface, command string, commandOptions ...string) {
	player := playerInterface.(*structs.Player)
	sendMessage := fmt.Sprintf("Inventory\r\n")
	for _, item := range player.Items {
		sendMessage += fmt.Sprintf("&y%s&n\r\n", item.GetName())
	}
	player.SendToPlayer(sendMessage)
}