package main

import (
	"velk/src/commands"
	"velk/src/infrastructure"
	"velk/src/interfaces"
	"velk/src/services"
	"velk/src/structs"
)

// only needed below for sample processing

func main() {
	infrastructure.CommandList = make(map[string]interfaces.CommandInterface)
	infrastructure.CommandList["say"] = &commands.Say{}
	infrastructure.CommandList["who"] = &commands.Who{}
	infrastructure.CommandList["shutdown"] = &commands.Shutdown{}
	infrastructure.CommandList["quit"] = &commands.Quit{}
	infrastructure.CommandList["look"] = &commands.Look{}
	infrastructure.CommandList["north"] = &commands.Move{}
	infrastructure.CommandList["east"] = &commands.Move{}
	infrastructure.CommandList["south"] = &commands.Move{}
	infrastructure.CommandList["west"] = &commands.Move{}

	infrastructure.Server = &structs.Server{}
	serviceService := services.ServerService{}
	serviceService.Init()
	go serviceService.CommandConsumer()
	serviceService.ConnectionLoop()

}