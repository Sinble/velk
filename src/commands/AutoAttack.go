package commands

import (
	"fmt"
	"velk/src/services"
	"velk/src/structs"
)

type AutoAttack struct {
	PlayerService services.PlayerService
}

func (a *AutoAttack) Action(player *structs.Player, command string, commandOptions ...string) {
	player.SendToPlayer(fmt.Sprintf("You punch %s\r\n", player.Targets[0].Name))
	player.State = "STANDING"
}