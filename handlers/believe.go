package handlers

import (
	"github.com/golang/glog"
	"code.google.com/p/goprotobuf/proto"
	"github.com/gysan/room/common"
	"github.com/gysan/room/packet"
	"github.com/gysan/room/game"
	"math/rand"
	"encoding/json"
	"time"
)

var (
	BelieveRoomManage game.RoomManage
	BelieveUserManage game.UserManage
)

/**
 * 30105
 */
func BelieveStartGame(roomId string, receivePacket *packet.Packet, message *common.Message) {
	glog.Infof("BelieveStartGame...........................................")
	room := BelieveRoomManage.GetRoom(roomId)
	if room == nil {
		glog.Errorf("The room [%v] is not exist.", roomId)
		return
	}

	users := room.GetUsers()
	length := len(users)
	glog.Infof("length: %v", length)
	tempUsers := make([]string, 0, length-1)
	for userId, _ := range room.GetUsers() {
		glog.Infof("The room has user: %v", userId)
		tempUsers = append(tempUsers, userId)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randIndex := r.Intn(length)
	room.SetGuessUser(tempUsers[randIndex])
	glog.Infof("RandIndex: %v Guess user: %v", randIndex, room.GetGuessUser())

	answerTempUsers := make([]string, 0, length-2)
	answerTempUsers = append(answerTempUsers, tempUsers[:randIndex]...)
	glog.Infof("answerTempUsers: %v", answerTempUsers)
	answerTempUsers = append(answerTempUsers, tempUsers[randIndex+1:]...)

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	randIndex = r.Intn(length-1)
	room.SetAnswerUser(answerTempUsers[randIndex])
	glog.Infof("RandIndex: %v Answer user: %v", randIndex, room.GetAnswerUser())
	glog.Infof("tempUsers: %v answerTempUsers: %v Room: %v", tempUsers, answerTempUsers, room)

	bodyJson := map[string]interface{}{
		"question" : "question",
		"answer" : "answer",
	}

	usersJson := make([]interface{}, 0, length-1)
	for i, userId := range answerTempUsers {
		userJson := map[string]interface{}{
			"user_id" : userId,
			"index" : i + 1,
			"guess" : 0,
			"status" : 0,
		}
		usersJson = append(usersJson, userJson)
	}
	usersJson = append(usersJson, map[string]interface{}{
			"user_id" : room.GetGuessUser(),
			"index" : length,
			"guess" : 1,
			"status" : 0,
		})
	bodyJson["users"] = usersJson

	body, err := json.Marshal(bodyJson)
	if err != nil {
		glog.Errorf("json.Marshal(\"%v\") error(%v)", bodyJson, err)
		return
	}
	glog.Infof("Body: %s", string(body))
	message.Body = proto.String(string(body))
	for _, userId := range tempUsers {
		glog.Infof("Send message to user: %v", userId)
		SendMessage(userId, receivePacket.Tag, message)
	}
	glog.Infof("BelieveStartGame end.......................................")
	return
}

/**
 * 30101
 */
func BelieveEnterRoom(roomId string, receivePacket *packet.Packet, message *common.Message) {
	glog.Infof("BelieveEnterRoom...........................................")
	sender := message.GetSender()
	room := BelieveRoomManage.GetRoom(roomId)
	if room == nil {
		room = BelieveRoomManage.CreateRoom(sender, roomId)
	}
	user := BelieveUserManage.CreateUser(sender)
	status := room.AddUser(user)
	switch status{
	case 0:
		glog.Infof("Room is full.")
		message.Receiver = proto.String(sender)
		message.Sender = proto.String(roomId)
		message.Type = proto.String("30102")
		SendMessage(sender, receivePacket.Tag, message)
	case 1:
	for userId, _ := range room.GetUsers() {
		glog.Infof("The room has user: %v", userId)
		SendMessage(userId, receivePacket.Tag, message)
	}
	}
	glog.Infof("room: %v user: %v roomManage:%v userManage:%v", room, user, BelieveRoomManage, BelieveUserManage)
	glog.Infof("BelieveEnterRoom end........................................")
	return
}

/**
 * 30103
 */
func BelieveOuterRoom(roomId string, receivePacket *packet.Packet, message *common.Message) {
	glog.Infof("BelieveOuterRoom............................................")
	sender := message.GetSender()

	room := BelieveRoomManage.GetRoom(roomId)
	if room == nil {
		glog.Errorf("The room [%v] is not exist.", roomId)
		return
	}

	user := BelieveUserManage.GetUser(sender)
	room.RemoveUser(user)
	BelieveUserManage.DeleteUser(user)

	if len(room.GetUsers()) > 0 {
		for userId, _ := range room.GetUsers() {
			glog.Infof("The room has user [%v]", userId)
			SendMessage(userId, receivePacket.Tag, message)
		}
	}else {
		BelieveRoomManage.DeleteRoom(room)
	}
	glog.Infof("room: %v user: %v roomManage:%v userManage:%v", room, user, BelieveRoomManage, BelieveUserManage)
	glog.Infof("BelieveOuterRoom end........................................")
	return
}

func BelieveRoom(roomId string, receivePacket *packet.Packet, message *common.Message) {
	glog.Infof("BelieveRoom............................................")
	//	sender := message.GetSender()

	room := BelieveRoomManage.GetRoom(roomId)
	if room == nil {
		glog.Errorf("The room [%v] is not exist.", roomId)
	}

	if len(room.GetUsers()) > 0 {
		for userId, _ := range room.GetUsers() {
			glog.Infof("The room has user [%v]", userId)
			SendMessage(userId, receivePacket.Tag, message)
		}
	}else {
		BelieveRoomManage.DeleteRoom(room)
	}
	glog.Infof("BelieveRoom end........................................")
	return
}
