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

func FindOfflineMessages(userId string) ([]*common.Message, error) {
	userIdInt, _ := strconv.Atoi(userId)
	glog.Infof("Find offline messages: %v", userIdInt%16)
	rows, err := db.GetDb().Query("select id, sender, receiver, body from message_"+strconv.Itoa(userIdInt%16)+" where sender = ? and status = 0 order by ctime asc", userId)
	if err != nil {
		glog.Errorf("db.Query(\"%s\") failed (%v)", "select id, sender, receiver, body from message_"+strconv.Itoa(userIdInt%16)+" where sender = ? and status = 0 order by ctime asc", err)
		return nil, err
	}
	var id, sender, receiver, body string
	messages := []*common.Message{}
	for rows.Next() {
		if err := rows.Scan(&id, &receiver, &sender, &body); err != nil {
			glog.Errorf("rows.Scan() failed (%v)", err)
			return nil, err
		}
		glog.Infof("%v", id)
		messages = append(messages, &common.Message{MessageId: proto.String(id), Sender: proto.String(sender), Receiver: proto.String(receiver), Body: proto.String(body)})
	}
	return messages, nil
}

func UpdateMessageStatus(receiver interface{}, messageId string, status int) error {
	receiverInt, _ := strconv.Atoi(receiver.(string))
	glog.Infof("Update status: %v, %v", receiverInt%16, status)
	timestamp := convert.TimestampToTimeString(time.Now().Unix())
	_, err := db.GetDb().Exec("update message_"+strconv.Itoa(receiverInt%16)+" set status = ?, utime = ? where id = ? and sender = ?", status, timestamp, messageId, receiver.(string))
	if err != nil {
		glog.Errorf("db.Exec(\"update message_"+strconv.Itoa(receiverInt%16)+" set status = ?, utime = ? where id = ? and sender = ?\") failed (%v)", err)
		return err
	}
	return nil
}

func InsertMessage(message *common.Message) error {
	senderInt, _ := strconv.Atoi(message.GetSender())
	receiverInt, _ := strconv.Atoi(message.GetReceiver())
	glog.Infof("Insert message: [sender: %v] [receiver: %v]", senderInt%16, receiverInt%16)
	timestamp := convert.TimestampToTimeString(message.GetDate() / 1000)
	_, err := db.GetDb().Exec("INSERT INTO `message_"+strconv.Itoa(senderInt%16)+"` (`id`, `sender`, `receiver`, `body`, `status`, `utime`, `ctime`) VALUES (?, ?, ?, ?, ?, ?, ?)",
		message.GetMessageId(), message.GetSender(), message.GetReceiver(), message.GetBody(), 10, timestamp, timestamp)
	_, senderErr := db.GetDb().Exec("INSERT INTO `message_"+strconv.Itoa(receiverInt%16)+"` (`id`, `sender`, `receiver`, `body`, `status`, `utime`, `ctime`) VALUES (?, ?, ?, ?, ?, ?, ?)",
		message.GetMessageId(), message.GetReceiver(), message.GetSender(), message.GetBody(), 0, timestamp, timestamp)
	if err != nil || senderErr != nil {
		glog.Errorf("db.Exec(\"INSERT INTO `message` (`id`, `sender`, `is_sender`, `receiver`, `body`, `utime`, `ctime`) VALUES (?, ?, ?, ?, ?, ?, ?)\") failed (%v)", err)
		return err
	}
	return nil
}
