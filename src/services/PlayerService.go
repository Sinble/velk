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



func (service PlayerService) SendToAllPlayers(message string) error {

	players, err := service.GetPlayers()
	if err != nil {
		return err
	}
	for _, player := range players {
		player.SendToPlayer(message)
	}
	return nil
}




