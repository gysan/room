package dao

import (
	"github.com/golang/glog"
	"github.com/gysan/room/utils/db"
)

func UpdateOnline(userId, status int) error {
	glog.Infof("Update status: %v, %v", userId, status)
	_, err := db.GetDb().Exec("update user set online = ? where id = ?", status, userId)
	if err != nil {
		glog.Errorf("db.Exec(\"update user set online = ? where id = ?\") failed (%v)", err)
		return err
	}
	return nil
}
