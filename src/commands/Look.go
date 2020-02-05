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
	sendMessage := fmt.Sprintf("&y%s&n\r\n", player.Room.Name)
	sendMessage += fmt.Sprintf("%s\r\n", player.Room.Description)
	sendMessage += "&cExits: [&y "
	if player.Room.NorthExit != nil {
		sendMessage += "N "
	}
	if player.Room.EastExit != nil {
		sendMessage += "E "
	}
	if player.Room.SouthExit != nil {
		sendMessage += "S "
	}
	if player.Room.WestExit != nil {
		sendMessage += "W "
	}
	sendMessage += "&c]&n\r\n"
	for _, player := range player.Room.GetPlayers() {
		sendMessage += fmt.Sprintf("&y%s&n\r\n", player.Name)
	}

	player.SendToPlayer(sendMessage)
}