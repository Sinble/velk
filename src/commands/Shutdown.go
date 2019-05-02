package commands

import (
	"velk/src/infrastructure"
	"velk/src/services"
	"velk/src/structs"
)

type Shutdown struct {
	PlayerService services.PlayerService
}

func (shutdown *Shutdown) Action(player *structs.Player, command string, commandOptions ...string) {

	shutdown.PlayerService.SendToAllPlayers("Velk Server Shutting Down\r\n")
	infrastructure.Server.Listener.Close()

}
