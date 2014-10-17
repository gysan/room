package game

import "github.com/golang/glog"

type RoomManage struct {
	rooms map[string]*Room
}

func (this *RoomManage) CreateRoom(userId string, roomId string) *Room {
	size := len(this.rooms)
	if size == 0 {
		this.rooms = make(map[string]*Room, 100)
	}else if size >= 100 {
		return nil
	}

	var room = new(Room)
	room.id = roomId
	room.admin = userId

	this.rooms[roomId] = room
	return room
}

func (this *RoomManage) DeleteRoom(room *Room) {
	if room == nil {
		return
	}
	glog.Infof("Delete the room [%v]", room.id)
	delete(this.rooms, room.id)
}

func (this *RoomManage) GetRooms() map[string]*Room {
	return this.rooms
}

func (this *RoomManage) GetRoom(id string) *Room {
	return this.rooms[id]
}
