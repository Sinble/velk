package infrastructure

import (
	"velk/src/interfaces"
	"velk/src/structs"
)

var Server *structs.Server

var Rooms map[int]structs.Room
var Players  map[int]structs.Player
var CommandList map[string]interfaces.CommandInterface


