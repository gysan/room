package game

import (
	"sync"
	"github.com/golang/glog"
)

type UserManage struct {
	users map[string]*User
	lock sync.Mutex
}

func (this *UserManage) CreateUser(userId string) *User{
	var user = new(User)
	user.id = userId

	if len(this.users) == 0 {
		this.users = make(map[string]*User)
	}
	this.lock.Lock()
	this.users[userId] = user
	this.lock.Unlock()
	return user
}

func (this *UserManage) GetUser(id string) *User{
	return this.users[id]
}

func (this *UserManage) DeleteUser(user *User){
	this.lock.Lock()
	delete(this.users, user.id)
	this.lock.Unlock()
	glog.Infof("The user [%v] is deleted.", user.id)
}
