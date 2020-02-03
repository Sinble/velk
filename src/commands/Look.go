package commands

import (
	"fmt"
	"velk/src/services"
	"velk/src/structs"
)

type Look struct {
	PlayerService services.PlayerService
}

func (look *Look) Action(player *structs.Player, command string, commandOptions ...string) {
	sendMessage := fmt.Sprintf("%s\r\n", player.Room.Name)
	sendMessage += fmt.Sprintf("%s\r\n", player.Room.Description)
	for _, player := range player.Room.Players {
		sendMessage += fmt.Sprintf("%s\r\n", player.Name)
	}

	look.PlayerService.SendToPlayer(player, sendMessage)
}
