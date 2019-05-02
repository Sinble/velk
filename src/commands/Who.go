package commands

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"velk/src/services"
	"velk/src/structs"
)

type Who struct {
	PlayerService services.PlayerService
}

func (who *Who) Action(player *structs.Player, command string, commandOptions ...string) {
	players, err := who.PlayerService.GetPlayers()
	if err != nil {
		logrus.Error(err)
	}
	sendMessage := "------------------------Velk--------------------------------\r\n"
	for _, player := range players {
		sendMessage += fmt.Sprintf("%s\r\n", player.Name)
	}
	sendMessage += "------------------------------------------------------------\r\n"
	sendMessage += fmt.Sprintf("Total Players: %d\r\n", len(players))
	who.PlayerService.SendToPlayer(player, sendMessage)
}

