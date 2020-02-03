package infrastructure

import (
	"velk/src/interfaces"
	"velk/src/structs"
)

var Server *structs.Server

var Zones map[int]*structs.Zone
var PlayerCount int
var Players  map[int]*structs.Player
var CommandList map[string]interfaces.CommandInterface

func init() {
	PlayerCount = 0
	Players = make(map[int]*structs.Player)
	Zones = make(map[int]*structs.Zone)
}

