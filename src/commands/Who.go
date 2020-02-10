package commands

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"velk/src/interfaces"
	"velk/src/services"
	"velk/src/structs"
)

type Who struct {
	PlayerService services.PlayerService
}

func (a *Who) Action(playerInterface interfaces.PlayerInterface, command string, commandOptions ...string) {
	player := playerInterface.(*structs.Player)
	players, err := a.PlayerService.GetPlayers()
	if err != nil {
		logrus.Error(err)
	}
	sendMessage := "------------------------Velk--------------------------------\r\n"
	for _, player := range players {
		sendMessage += fmt.Sprintf("%s\r\n", player.Name)
	}
	sendMessage += "------------------------------------------------------------\r\n"
	sendMessage += fmt.Sprintf("Total players: %d\r\n", len(players))
	player.SendToPlayer(sendMessage)
}

