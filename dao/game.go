package dao

import (
	"github.com/gysan/room/common"
	"strconv"
	"github.com/golang/glog"
	"github.com/gysan/room/utils/convert"
	"github.com/gysan/room/utils/db"
	"time"
)

func UpdateGameMessageStatus(receiver interface{}, messageId string, status int) error {
	receiverInt, _ := strconv.Atoi(receiver.(string))
	glog.Infof("Update status: %v, %v", receiverInt%16, status)
	timestamp := convert.TimestampToTimeString(time.Now().Unix())
	_, err := db.GetDb().Exec("update game_"+strconv.Itoa(receiverInt%16)+" set status = ?, utime = ? where id = ? and sender = ?", status, timestamp, messageId, receiver.(string))
	if err != nil {
		glog.Errorf("db.Exec(\"update game_"+strconv.Itoa(receiverInt%16)+" set status = ?, utime = ? where id = ? and sender = ?\") failed (%v)", err)
		return err
	}
	return nil
}

func InsertGameMessage(message *common.Message) error {
	senderInt, _ := strconv.Atoi(message.GetSender())
	receiverInt, _ := strconv.Atoi(message.GetReceiver())
	typeInt, _ := strconv.Atoi(message.GetType())
	glog.Infof("Insert game: [sender: %v] [receiver: %v]", senderInt%16, receiverInt%16)
	timestamp := convert.TimestampToTimeString(message.GetDate() / 1000)
	_, err := db.GetDb().Exec("INSERT INTO `game_"+strconv.Itoa(senderInt%16)+"` (`id`, `sender`, `receiver`, `type`, `body`, `status`, `utime`, `ctime`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		message.GetMessageId(), message.GetSender(), message.GetReceiver(), typeInt, message.GetBody(), 10, timestamp, timestamp)
	_, senderErr := db.GetDb().Exec("INSERT INTO `game_"+strconv.Itoa(receiverInt%16)+"` (`id`, `sender`, `receiver`, `type`, `body`, `status`, `utime`, `ctime`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		message.GetMessageId(), message.GetReceiver(), message.GetSender(), typeInt, message.GetBody(), 0, timestamp, timestamp)
	if err != nil || senderErr != nil {
		glog.Errorf("db.Exec(\"INSERT INTO `game` (`id`, `sender`, `is_sender`, `receiver`, `type`,  `body`, `utime`, `ctime`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)\") failed (%v)", err)
		return err
	}
	return nil
}
