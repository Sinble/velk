package services

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"strings"
	"velk/src/infrastructure"
	"velk/src/interfaces"
	"velk/src/repositories"
	"velk/src/structs"
)

type ServerService struct {
	PlayerService PlayerService
	WorldRepository repositories.WorldRepository
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

	zone := serverService.WorldRepository.CreateZone("Sinble")
	room1 := serverService.WorldRepository.CreateRoom(zone.Id)
	room2 := serverService.WorldRepository.CreateRoom(zone.Id)
	room1.NorthExit = room2
	room2.SouthExit = room1
	room1.AddMob(&structs.Mob{
		Id:    0,
		Name:  "Joe",
		Room:  room1,
		State: "STANDING",
		Health: 0,
	})

	room1.Items[0] = &structs.Item{
		ID: 0,
		Name: "sword",
	}

}

func (serverService *ServerService) ConnectionLoop() {
	for {
		connection, err := infrastructure.Server.Listener.Accept()

		if err != nil {
			logrus.Error(err)
			return
		}
		player := &structs.Player{Connection: connection, Reader:bufio.NewReader(connection), FightQueue: make(chan structs.Command), FightChannel: make(chan bool), Items: make(map[int]interfaces.ItemInterface)}

		go serverService.PlayerGameLoop(player)
	}
}

func (serverService *ServerService) PlayerGameLoop(player *structs.Player) {

	player.SendToPlayer("Welcome to Velk\r\nWhat is thy name?")
	name, err := player.ReadFromPlayer()

	if err != nil {
		logrus.Error("Failed to read name", err)
		return
	}

	player.Name = name
	player.Health = 10
	player.MaxHealth = 10
	serverService.PlayerService.SendToAllPlayers(fmt.Sprintf("%s has joined the server\r\n", player.Name))
	serverService.PlayerService.AddPlayer(player)
	go serverService.PlayerService.waitFight(player)
	room, err := serverService.WorldRepository.GetRoom("1-1")
	if err !=nil {
		logrus.Error(err)
	}
	player.Room = room
	room.AddPlayer(player)
	//commands.LookAtRoom(player)
	player.SendPlayerPrompt()
	//Room.AddPlayer(player)
	for {
		commandString, err := player.ReadFromPlayer()

		if err != nil {
			logrus.Error("Failed to receive commandName: ", err)
			player.Connection.Close()
			serverService.PlayerService.RemovePlayer(player.Id)
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
			command.Player.SendToPlayer("Huh?\r\n")
		}

		command.Player.SendPlayerPrompt()

	}
}

func (serverService *ServerService) processCommandString(commandString string) (string, string) {

	commandSplit := strings.SplitN(commandString, " ", 2)

	if len(commandSplit) < 2 {
		return commandSplit[0], ""
	}

	return strings.ToLower(commandSplit[0]), commandSplit[1]
}
