package commands

import (
	"velk/src/infrastructure"
	"velk/src/interfaces"
	"velk/src/services"
)

type Shutdown struct {
	PlayerService services.PlayerService
}

func (shutdown *Shutdown) Action(playerInterface interfaces.PlayerInterface, command string, commandOptions ...string) {

	shutdown.PlayerService.SendToAllPlayers("Velk Server Shutting Down\r\n")
	infrastructure.Server.Listener.Close()

}
