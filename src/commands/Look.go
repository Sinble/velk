package commands

import (
	"fmt"
	"strings"
	"velk/src/interfaces"
	"velk/src/services"
	"velk/src/structs"
)

type Look struct {
	PlayerService services.PlayerService
}

func (a *Look) Action(playerInterface interfaces.PlayerInterface, command string, commandOptions ...string) {
	player := playerInterface.(*structs.Player)

	options := strings.Split(commandOptions[0], " ")
	if len(options) > 1 {
		if strings.ToLower(options[0]) == "in" {
			for _, item := range player.Room.Items {
				if item.GetName() == options[1] {
					container := item.(*structs.Container)
					sendMessage := "Items in \r\n"
					for _, containerItem := range container.Items {
						sendMessage += fmt.Sprintf("%s\r\n",  containerItem.GetName())
					}
					player.SendToPlayer(sendMessage)
				}
			}
		}
	}
	if commandOptions[0] != "" {
		for id, mob := range player.Room.GetMobs() {
			if strings.Contains(mob.Name, commandOptions[0]) {
				mob, lock := player.Room.MobCheck(id)
				if mob != nil {
					sendMessage:= "blah\r\n"
					player.SendToPlayer(sendMessage)
					lock.Unlock()
					return
				}
			}
		}
		player.SendToPlayer("I dont see that\r\n")
	} else {
		LookAtRoom(player)
	}
}

func LookAtRoom(player *structs.Player) {
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
	for _, item := range player.Room.Items {
		sendMessage += fmt.Sprintf("&y%s&n\r\n", item.GetName())
	}

	for _, player := range player.Room.GetPlayers() {
		sendMessage += fmt.Sprintf("&y%s&n\r\n", player.Name)
	}
	for _, mob := range player.Room.GetMobs() {
		sendMessage += fmt.Sprintf("&y%s&n\r\n", mob.Name)
	}


	player.SendToPlayer(sendMessage)
}