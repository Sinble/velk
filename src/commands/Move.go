package commands

import (
	"velk/src/services"
	"velk/src/structs"
)

type Move struct {
	PlayerService services.PlayerService
}

func (move *Move) Action(player *structs.Player, command string, commandOptions ...string) {
	if player.State == "FIGHTING" {
		player.SendToPlayer("How are you suppose to do that your fighting!\r\n")
		return
	}
	switch command {
	case "north":
		if player.Room.NorthExit != nil {
			player.Room.RemovePlayer(player)
			player.Room.NorthExit.AddPlayer(player)
			player.Room = player.Room.NorthExit
			LookAtRoom(player)
		} else {
			player.SendToPlayer("You smash your face into a wall\r\n")
		}
		break
	case "east":
		if player.Room.EastExit != nil {
			player.Room.RemovePlayer(player)
			player.Room.EastExit.AddPlayer(player)
			player.Room = player.Room.EastExit
			LookAtRoom(player)
		} else {
			player.SendToPlayer("You smash your face into a wall\r\n")
		}
		break
	case "south":
		if player.Room.SouthExit != nil {
			player.Room.RemovePlayer(player)
			player.Room.SouthExit.AddPlayer(player)
			player.Room = player.Room.SouthExit
			LookAtRoom(player)
		} else {
			player.SendToPlayer("You smash your face into a wall\r\n")
		}
		break
	case "west":
		if player.Room.WestExit != nil {
			player.Room.RemovePlayer(player)
			player.Room.WestExit.AddPlayer(player)
			player.Room = player.Room.WestExit
			LookAtRoom(player)
		} else {
			player.SendToPlayer("You smash your face into a wall\r\n")
		}
		break
	default:
		player.SendToPlayer("")
	}
}
