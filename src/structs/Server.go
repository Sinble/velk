package structs

import "net"

type Server struct {
	Listener net.Listener
	PlayerCommandChannel chan Command
}