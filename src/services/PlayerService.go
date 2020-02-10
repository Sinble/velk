package services

import (
	"time"
	"velk/src/infrastructure"
	"velk/src/repositories"
	"velk/src/structs"

	"github.com/sirupsen/logrus"
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
		player.SendToPlayer( message)
	}
	return nil
}

func(service PlayerService) waitFight(player *structs.Player) {
	logrus.Debug("here")
	for range player.FightChannel {
		logrus.Debug("here1")
		go service.mobAttack(player.Targets[0])
		service.Fighting(player)
	}
}

func (service PlayerService) mobAttack(mob *structs.Mob) {
	duration := 1
	for true {
		infrastructure.Server.PlayerCommandChannel <- structs.Command{Player: mob, CommandName:"autoattack", CommandSuffix: ""}
		time.Sleep(time.Duration(duration) * time.Second)
		if mob.State != "FIGHTING" {
			break
		}
	}
}
func(service PlayerService) Fighting(player *structs.Player) {

	duration := 1
	for true {
		select {
		case command := <-player.FightQueue:
			infrastructure.Server.PlayerCommandChannel <- command
		default:
			infrastructure.Server.PlayerCommandChannel <- structs.Command{Player: player, CommandName:"autoattack", CommandSuffix: ""}
		}
		
		time.Sleep(time.Duration(duration) * time.Second)
		if player.State != "FIGHTING" {
			break
		}
	}
}




