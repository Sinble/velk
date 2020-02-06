package structs

import (
	"bufio"
	"fmt"
	"net"
	"velk/src/utils"

	"github.com/sirupsen/logrus"
)

type Player struct {
	Id int
	Name string
	Connection net.Conn
	Reader *bufio.Reader
	Room *Room
	State string
	Targets []*Mob

	Health int
	MaxHealth int
}

func (p Player) SendToPlayer(message string) {
	cs := utils.ColorService{}.New()
	message = cs.ProcessString(message)
	_, err := p.Connection.Write([]byte(message))
	if err != nil {
		logrus.Errorf("Player %s: %s", p.Name, err)
	}
}

func (p Player) ReadFromPlayer() (string, error) {
	message, err := p.Reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return message[:len(message)-2], nil
}

func(p Player) SendPlayerPrompt() {
	p.SendToPlayer(fmt.Sprintf("Health: %d/%d >", p.Health, p.MaxHealth))
}