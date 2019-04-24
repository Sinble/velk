package services

import (
	"velk/src/structs"
)

type PlayerService struct {

}

func (playerService PlayerService) SendToPlayer(player structs.Player, message string) error {
	_, err := player.Connection.Write([]byte(message))

	return err
}

func (playerService PlayerService) ReadFromPlayer(player structs.Player) (string, error) {
	message, err := player.Reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return message[:len(message)-2], nil
}

