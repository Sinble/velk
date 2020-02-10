package commands

import (
	"fmt"
	"velk/src/interfaces"
	"velk/src/services"
	"velk/src/structs"
)

type Quit struct {
	PlayerService services.PlayerService
}

func (quit *Quit) Action(playerInterface interfaces.PlayerInterface, command string, commandOptions ...string) {
	player := playerInterface.(*structs.Player)
	playerName := player.Name
	//TODO move this to playerservice
	player.Connection.Close()
	quit.PlayerService.SendToAllPlayers(fmt.Sprintf("%s has left the server\r\n", playerName))


}
