package room

import "github.com/golang/glog"

type User struct {
	id string
	room *Room
}

func (this *User) GetRoom() *Room{
	return this.room
}

func (this *User) CreateRoom(id string){
	glog.Infof("Create room: %v", id)
}

func (this *User) JoinRoom(id string){
	glog.Infof("Join room: %v", id)
}
