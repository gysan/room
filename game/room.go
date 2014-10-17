package game

import (
	"sync"
	"github.com/gysan/room/config"
	"github.com/golang/glog"
)

type Room struct {
	id             string
	admin          string
	password       string
	question       string
	answer         string
	users          map[string]*User
	guessUser      string
	answerUser     string
	eliminateUsers map[string]*User
	outUsers       map[string]*User
	lock           sync.Mutex
}

func (this *Room) AddUser(user *User) int {
	if user == nil {
		return -1
	}

	this.lock.Lock()
	count := len(this.users)
	if count >= config.BelieveRoomUserLimit {
		glog.Infof("The room is full of user. The user [%v] is denied", user.id)
		this.lock.Unlock()
		return 0
	}

	user.room = this
	if count == 0 {
		this.users = make(map[string]*User, config.BelieveRoomUserLimit)
	}
	this.users[user.id] = user
	this.lock.Unlock()
	glog.Infof("The user [%v] enters the room.", user.id)
	return 1
}

func (this *Room) RemoveUser(user *User) {
	glog.Infof("The user [%v] is removed from the room [%v].", user.id, this.id)

	if this.admin == user.id {
		glog.Infof("The user [%v] is admin.", user.id)
	}

	this.lock.Lock()
	delete(this.users, user.id)
	this.lock.Unlock()
}

func (this *Room) GetUsers() map[string]*User {
	return this.users
}

func (this *Room) SetGuessUser(user string) {
	this.guessUser = user
}

func (this *Room) GetGuessUser() string {
	return this.guessUser
}
func (this *Room) SetAnswerUser(user string) {
	this.answerUser = user
}

func (this *Room) GetAnswerUser() string {
	return this.answerUser
}


