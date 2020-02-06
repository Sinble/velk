package commands

import (
	"strings"
	"velk/src/services"
	"velk/src/structs"
)

type Kill struct {
	PlayerService services.PlayerService
}

func (a *Kill) Action(player *structs.Player, command string, commandOptions ...string) {
	if commandOptions[0] != "" {
		for id, mob := range player.Room.GetMobs() {
			if strings.Contains(mob.Name, commandOptions[0]) {
				mob, lock := player.Room.MobCheck(id)
				if mob != nil {
					player.State = "FIGHTING"
					player.Targets = append(player.Targets, mob)
					mob.Targets = append(mob.Targets, player)
					mob.State = "FIGHTING"
					lock.Unlock()
					return
				}
			}
		}
		player.SendToPlayer("I don't see them\r\n")
	} else {
		player.SendToPlayer("Who do you want me to kill???\r\n")
	}
}
