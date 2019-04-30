package services

import (
	"velk/src/repositories"
	"velk/src/structs"
)

type PlayerService struct {
	PlayerRepository repositories.PlayerRepository

}

func (service PlayerService) GetPlayer(id int) (*structs.Player, error) {
	return service.PlayerRepository.GetPlayer(id)
}

func (service PlayerService) RemovePlayer(id int) error {
	return service.PlayerRepository.RemovePlayer(id)
}

func (service PlayerService) GetPlayers() ([]*structs.Player, error) {
	return service.PlayerRepository.GetPlayers()
}

func (service PlayerService) AddPlayer(player *structs.Player) error {
	return service.PlayerRepository.AddPlayer(player)
}

func (service PlayerService) SendToPlayer(player *structs.Player, message string) error {
	_, err := player.Connection.Write([]byte(message))

	return err
}

func (service PlayerService) ReadFromPlayer(player *structs.Player) (string, error) {
	message, err := player.Reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return message[:len(message)-2], nil
}

