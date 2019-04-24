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
	Room Room
}
