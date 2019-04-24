package repositories

import (
	"velk/src/infrastructure"
	"velk/src/structs"
)

type PlayerRepository struct {

}

func (repository *PlayerRepository) GetPlayers() ([]structs.Player, error){
	var players []structs.Player

	for _, value := range infrastructure.Players {
		players = append(players, value)
	}

	return players, nil
}

func (repository *PlayerRepository) GetPlayer(id int) (structs.Player, error) {
	player, playerExists := infrastructure.Players[id]

	if !playerExists {

	}

	return player, nil
}



func (repository *PlayerRepository) AddPlayer(player structs.Player) error {
	infrastructure.Players[player.Id] = player

	return nil
}

func (repository *PlayerRepository) RemovePlayer(id int) error {
	delete(infrastructure.Players, id)
	return nil
}