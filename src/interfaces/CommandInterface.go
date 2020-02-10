package interfaces

type CommandInterface interface {
	Action(PlayerInterface, string, ...string)
}
