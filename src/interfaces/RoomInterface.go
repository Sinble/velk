package interfaces

type RoomInterface interface {
	AddPlayer(playerInterface PlayerInterface)
	RemovePlayer(playerInterface PlayerInterface)
	GetPlayers() []PlayerInterface
}
