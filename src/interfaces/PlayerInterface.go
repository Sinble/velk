package interfaces

type PlayerInterface interface {
	SendToPlayer(string)
	ReadFromPlayer() (string, error)
	SendPlayerPrompt()
}
