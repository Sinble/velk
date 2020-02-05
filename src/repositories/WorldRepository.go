package repositories

import (
	"errors"
	"strconv"
	"strings"
	"velk/src/infrastructure"
	"velk/src/structs"
)

type WorldRepository struct {

}

//func (repository *WorldRepository) GetRooms() ([]*structs.Room, error){
//	var players []*structs.Room
//
//	for _, value := range infrastructure.Rooms {
//		players = append(players, value)
//	}
//
//	return players, nil
//}

func (repository *WorldRepository) CreateZone(creator string) *structs.Zone {
	zoneId := len(infrastructure.Zones) + 1
	zone := &structs.Zone{
		Id:      zoneId,
		Name:    "New Zone",
		Creator: creator,
		Rooms:   make(map[int]*structs.Room),
	}
	infrastructure.Zones[zoneId] = zone

	return zone
}

func (repository *WorldRepository) CreateRoom(zoneId int) *structs.Room {
	roomID := len(infrastructure.Zones[zoneId].Rooms) + 1
	room := structs.Room{}.New(roomID)
	infrastructure.Zones[zoneId].Rooms[roomID] = room
	return room
}

func (repository *WorldRepository) GetRoom(id string) (*structs.Room, error) {
	split := strings.Split(id, "-")
	zoneStr := split[0]
	roomStr := split[1]

	zoneId, err := strconv.Atoi(zoneStr)
	if err != nil {
		return nil, err
	}

	roomId, err := strconv.Atoi(roomStr)
	if err != nil {
		return nil, err
	}

	zone, zoneExists := infrastructure.Zones[zoneId]

	if !zoneExists {
		return nil, errors.New("zone does not exist")
	}

	room, roomExists := zone.Rooms[roomId]

	if !roomExists {
		return nil, errors.New("room not exist in zone")
	}

	return room, nil
}



//func (repository *WorldRepository) AddRoom(player *structs.Room) error {
//
//	player.Id = infrastructure.PlayerCount
//	infrastructure.PlayerCount++
//	infrastructure.players[player.Id] = player
//
//	return nil
//}
//
//func (repository *WorldRepository) RemoveRoom(id int) error {
//	delete(infrastructure.Rooms, id)
//	return nil
//}
