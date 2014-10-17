package dao

import (
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gysan/room/common"
	"code.google.com/p/goprotobuf/proto"
	"time"
	"github.com/gysan/room/utils/convert"
	"strconv"
	"github.com/gysan/room/utils/db"
)

func FindOfflineNormalMessages(userId string) ([]*common.NormalMessage, error) {
	userIdInt, _ := strconv.Atoi(userId)
	glog.Infof("Find offline messages: %v", userIdInt%16)
	rows, err := db.GetDb().Query("select id, receiver, body, ctime from normal_message_"+strconv.Itoa(userIdInt%16)+" where receiver = ? and status = 0 order by ctime asc", userId)
	if err != nil {
		glog.Errorf("db.Query(\"%s\") failed (%v)", "select id, receiver, body, ctime from normal_message_"+strconv.Itoa(userIdInt%16)+" where receiver = ? and status = 0 order by ctime asc", err)
		return nil, err
	}
	var id, receiver, body string
	var date time.Time
	messages := []*common.NormalMessage{}
	for rows.Next() {
		if err := rows.Scan(&id, &receiver, &body, &date); err != nil {
			glog.Errorf("rows.Scan() failed (%v)", err)
			return nil, err
		}
		glog.Infof("%v", id)
		messages = append(messages, &common.NormalMessage{MessageId: proto.String(id), Receiver: proto.String(receiver), Content: []byte(body), Date: proto.Int64(date.Unix())})
	}
	return messages, nil
}

func UpdateNormalMessageStatus(receiver interface{}, messageId string, status int) error {
	receiverInt, _ := strconv.Atoi(receiver.(string))
	glog.Infof("Update status: %v, %v", receiverInt%16, status)
	timestamp := convert.TimestampToTimeString(time.Now().Unix())
	_, err := db.GetDb().Exec("update normal_message_"+strconv.Itoa(receiverInt%16)+" set status = ?, utime = ? where id = ?", status, timestamp, messageId)
	if err != nil {
		glog.Errorf("db.Exec(\"update normal_message_"+strconv.Itoa(receiverInt%16)+" set status = ?, utime = ? where id = ?\") failed (%v)", err)
		return err
	}
	return nil
}

func InsertNormalMessage(message *common.NormalMessage) error {
	receiverInt, _ := strconv.Atoi(message.GetReceiver())
	glog.Infof("Insert message: [receiver: %v]", receiverInt%16)
	timestamp := convert.TimestampToTimeString(time.Now().Unix())
	_, err := db.GetDb().Exec("INSERT INTO `normal_message_"+strconv.Itoa(receiverInt%16)+"` (`id`, `receiver`, `body`, `status`, `utime`, `ctime`) VALUES (?, ?, ?, ?, ?, ?)",
		message.GetMessageId(), message.GetReceiver(), message.GetContent(), 0, timestamp, timestamp)
	if err != nil {
		glog.Errorf("db.Exec(\"INSERT INTO `message` (`id`, `receiver`, `body`, `utime`, `ctime`) VALUES (?, ?, ?, ?, ?, ?, ?)\") failed (%v)", err)
		return err
	}
	return nil
}
