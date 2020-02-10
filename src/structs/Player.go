package structs

import (
	"bufio"
	"net"
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