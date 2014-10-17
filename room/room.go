package room

import (
	"sync"
	"github.com/golang/glog"
)

type Room struct {
	id   string
	users map[string]*User
	lock sync.Mutex
}

func (this *Room) GetUsers() map[string]*User {
	return this.users
}

func (this *Room) AddUser(user *User) bool {
	if user == nil {
		return false
	}

	this.lock.Lock()
	count := len(this.users)
	if count >= 5 {
		glog.Infof("The room is max user: %v", count)
		this.lock.Unlock()
	}
	user.room = this
	if count == 0 {
		this.users = make(map[string]*User, 5)
	}
	this.users[user.id] = user
	this.lock.Unlock()
	glog.Infof("The user: %v is inter room: %v", user.id, this.id)
	return true
}

func (this *Room) RemoveUser(user *User) {
	glog.Infof("The user: %v outer room: %v", user.id, this.id)

	this.lock.Lock()
	delete(this.users, user.id)
	this.lock.Unlock()

	if len(this.users) == 0{
		go func() {
			// delete room
		}()
	}
}

func (this *Room) GetUser(id string) *User {
	return this.users[id]
}
