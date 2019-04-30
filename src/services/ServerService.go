package services

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"strings"
	"velk/src/infrastructure"
	"velk/src/structs"
)

type ServerService struct {
	PlayerService PlayerService
}


func (serverService *ServerService) Init() {
	fmt.Println("Launching serverService...")

	// listen on all interfaces

	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		logrus.Fatal("Failed to init serverService", err)
	}

	infrastructure.Server.Listener = listener
	infrastructure.Server.PlayerCommandChannel = make(chan structs.Command)

}

func (serverService *ServerService) ConnectionLoop() {
	for {
		connection, err := infrastructure.Server.Listener.Accept()

		if err != nil {
			fmt.Println()
		}
		player := &structs.Player{Connection: connection, Reader:bufio.NewReader(connection)}
		serverService.PlayerService.AddPlayer(player)

		go serverService.PlayerGameLoop(player)
	}
}

func (serverService *ServerService) PlayerGameLoop(player *structs.Player) {

	serverService.PlayerService.SendToPlayer(player, "Welcome to Velk\r\nWhat is thy name?")
	name, err := serverService.PlayerService.ReadFromPlayer(player)

	if err != nil {
		logrus.Error("Failed to read name", err)
	}

	player.Name = name

	//Room.AddPlayer(player)
	for {
		commandString, err := serverService.PlayerService.ReadFromPlayer(player)

		if err != nil {
			logrus.Error("Failed to receive commandName: ", err)
			player.Connection.Close()
			return
		}

		commandName, commandSuffix := serverService.processCommandString(commandString)
		infrastructure.Server.PlayerCommandChannel <- structs.Command{Player: player, CommandName:commandName, CommandSuffix: commandSuffix}

	}
}

func (serverService *ServerService) CommandConsumer() {

	for command := range infrastructure.Server.PlayerCommandChannel {

		action, actionExists := infrastructure.CommandList[command.CommandName]
		if actionExists {
			action.Action(command.Player, command.CommandName, command.CommandSuffix)
		} else {
			serverService.PlayerService.SendToPlayer(command.Player,"Huh?\r\n")
		}

	}
}

func (serverService *ServerService) processCommandString(commandString string) (string, string) {

	commandSplit := strings.SplitN(commandString, " ", 2)

	if len(commandSplit) < 2 {
		return commandSplit[0], ""
	}

	return strings.ToLower(commandSplit[0]), commandSplit[1]
}
