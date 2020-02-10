package commands

import (
	"fmt"
	"velk/src/interfaces"
	"velk/src/services"
	"velk/src/structs"
)

type AutoAttack struct {
	PlayerService services.PlayerService
}

func (a *AutoAttack) Action(playerInterface interfaces.PlayerInterface, command string, commandOptions ...string) {


	switch playerInterface.(type) {
	case *structs.Player:
		player := playerInterface.(*structs.Player)
		player.SendToPlayer(fmt.Sprintf("You punch %s\r\n", player.Targets[0].Name))
		player.Targets[0].Health -= 1
		if player.Targets[0].Health <= 0 {
			death(player.Targets[0])
			player.State = "STANDING"
		}

	case *structs.Mob:
		player := playerInterface.(*structs.Mob)
		player.Targets[0].SendToPlayer(fmt.Sprintf("%s punches you\r\n", player.Name))
		player.Targets[0].Health -= 1

		player.State = "STANDING"
	}
}

func death(mob *structs.Mob) {
	mob.State = "DEAD"
	mob.Targets = mob.Targets[:0]
	mob.Room.RemoveMob(mob)
	mob = nil
}