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

	infrastructure.Server = &structs.Server{}
	serviceService := services.ServerService{}
	serviceService.Init()
	go serviceService.ConnectionLoop()
	serviceService.CommandConsumer()


}