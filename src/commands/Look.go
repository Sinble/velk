package commands

import (
	"fmt"
	"velk/src/services"
	"velk/src/structs"
)

type Look struct {
	PlayerService services.PlayerService
}

func (a *Look) Action(player *structs.Player, command string, commandOptions ...string) {
	CmdLook(player)
}

func CmdLook(player *structs.Player) {
	sendMessage := fmt.Sprintf("%s\r\n", player.Room.Name)
	sendMessage += fmt.Sprintf("%s\r\n", player.Room.Description)
	for _, player := range player.Room.GetPlayers() {
		sendMessage += fmt.Sprintf("%s\r\n", player.Name)
	}

	player.SendToPlayer(sendMessage)
}