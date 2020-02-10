package structs

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"velk/src/utils"
)

type Player struct {
	Id int
	Name string
	Connection net.Conn
	Reader *bufio.Reader
	Room *Room
	State string
	Targets []*Mob
	FightChannel chan bool
	FightQueue chan Command

	Health int
	MaxHealth int
}

func (player Player) SendToPlayer(message string) {
	cs := utils.ColorService{}.New()
	message = cs.ProcessString(message)
	_, err := player.Connection.Write([]byte(message))
	if err != nil {
		logrus.Errorf("Player %s: %s", player.Name, err)
	}
}

func (player Player) ReadFromPlayer() (string, error) {
	message, err := player.Reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return message[:len(message)-2], nil
}

func(player Player) SendPlayerPrompt() {
	player.SendToPlayer(fmt.Sprintf("Health: %d/%d >", player.Health, player.MaxHealth))
}