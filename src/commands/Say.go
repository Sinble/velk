package commands

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"velk/src/services"
	"velk/src/structs"
)

type Say struct {
	PlayerService services.PlayerService
}

func (say *Say) Action(player *structs.Player, command string, commandOptions ...string) {
	message := ""
	if len(commandOptions) > 0 {
		message = commandOptions[0]
	}
	sendMessage := fmt.Sprintf("%s says, \"%s\"\r\n", player.Name, message)
	players, err := say.PlayerService.GetPlayers()
	if err != nil {
		logrus.Error(err)
	}
	say.SendToAllPlayers(players, sendMessage)
}

func (say *Say) SendToAllPlayers(players []*structs.Player, message string)  {

	for _, player := range players {
		say.PlayerService.SendToPlayer(player, message)
	}
}
