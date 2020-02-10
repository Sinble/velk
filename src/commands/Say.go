package commands

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"velk/src/interfaces"
	"velk/src/services"
	"velk/src/structs"
)

type Say struct {
	PlayerService services.PlayerService
}

func (say *Say) Action(playerInterface interfaces.PlayerInterface, command string, commandOptions ...string) {
	player := playerInterface.(*structs.Player)
	message := ""
	if len(commandOptions) > 0 {
		message = commandOptions[0]
	}
	sendMessage := fmt.Sprintf("%s says, \"%s&n\"\r\n", player.Name, message)

	err := say.PlayerService.SendToAllPlayers(sendMessage)
	if err != nil {
		logrus.Error(err)
	}
}

