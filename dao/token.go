package dao

import (
	"github.com/golang/glog"
	"time"
	"github.com/gysan/room/utils/db"
	"github.com/gysan/room/utils/convert"
	"strconv"
	"github.com/gysan/room/push"
)

func InsertToken(userId, token string) error {
	userIdInt, _ := strconv.Atoi(userId)
	glog.Infof("Insert token: [userId: %v]", userIdInt)
	timestamp := convert.TimestampToTimeString(time.Now().Unix())
	_, err := db.GetDb().Exec("INSERT INTO `token` (`user_id`, `token`, `ctime`) VALUES (?, ?, ?)",
		userIdInt, token, timestamp)
	if err != nil {
		glog.Errorf("db.Exec(\"INSERT INTO `token` (`user_id`, `token`, `ctime`) VALUES (?, ?, ?)\") failed (%v)", err)
		return err
	}
	return nil
}

func FindTokensByUserId(userId string) ([]*push.Message, error) {
	userIdInt, _ := strconv.Atoi(userId)
	rows, err := db.GetDb().Query("select token from token where user_id = ?", userIdInt)
	if err != nil {
		glog.Errorf("db.Query(\"%s\") failed (%v)", "select token from token where user_id = ?", err)
		return nil, err
	}
	messages := []*push.Message{}
	for rows.Next() {
		var token string
		if err := rows.Scan(&token); err != nil {
			glog.Errorf("rows.Scan() failed (%v)", err)
			return nil, err
		}
		glog.Infof("%v", token)
		messages = append(messages, &push.Message{Token: token})
	}
	return messages, nil
}
