package handlers

import (
	"github.com/gysan/room/common"
	"github.com/golang/glog"
	"github.com/gysan/room/dao"
	"net"
	"github.com/gysan/room/packet"
	"code.google.com/p/goprotobuf/proto"
	"time"
	"github.com/gysan/room/push"
)

func HandleMessage(conn *net.TCPConn, receivePacket *packet.Packet) {
	glog.Info("HandleMessage beginning...")
	message := &common.Message{}
	packet.Unpack(receivePacket, message)
	message.Date = proto.Int64(time.Now().Unix()*1000)

	receiver := message.GetReceiver()
	messageType := message.GetType()
	glog.Infof("Message type: [%v]", messageType)

	switch string(messageType[0]){
	case "3":
		dao.InsertGameMessage(message)
		glog.Infof("Send message to game room id: [%v]", receiver)
		switch messageType{
		case "30101":
			// 进入房间
			BelieveEnterRoom(receiver, receivePacket, message)
		case "30103":
			// 退出房间
			BelieveOuterRoom(receiver, receivePacket, message)
		case "30105":
			// 开始游戏
			BelieveStartGame(receiver, receivePacket, message)
		default:
			BelieveRoom(receiver, receivePacket, message)
		}
	case "2":
		glog.Infof("Send message to group room id: [%v]", receiver)
	default:
		dao.InsertMessage(message)
		glog.Infof("Message: %v string: %s", receivePacket, message.String())
		glog.Info("HandleMessage end.")

		glog.Info("MessageResponse...")
		messageResponse := &common.MessageResponse{
			MessageId: proto.String(message.GetMessageId()),
			Status: proto.Bool(true),
		}
		glog.Infof("MessageResponse: %v string: %s", messageResponse, messageResponse.String())
		err := SendPbData(conn, receivePacket.Tag+1, uint32(common.MessageCommand_MESSAGE_RESPONSE), messageResponse)
		if err != nil {
			glog.Errorf("Send data error: %v", err)
			return
		}
		glog.Info("MessageResponse end.")
		glog.Infof("Send message to receiver: [%v]", receiver)
		SendMessage(receiver, receivePacket.Tag, message)
	}
	glog.Info("Send message to receiver end.")
}

func HandleReceiveMessageAck(conn *net.TCPConn, receivePacket *packet.Packet) {
	glog.Info("HandleReceiveMessageAck beginning...")
	receiveMessageAck := &common.ReceiveMessageAck{}
	packet.Unpack(receivePacket, receiveMessageAck)
	messageType := receiveMessageAck.GetType()
	switch string(messageType[0]){
	case "3":
		dao.UpdateGameMessageStatus(ConnMapUser.Get(conn), receiveMessageAck.GetMessageId(), int(receiveMessageAck.GetStatus()))
	case "2":

	default:
		dao.UpdateMessageStatus(ConnMapUser.Get(conn), receiveMessageAck.GetMessageId(), int(receiveMessageAck.GetStatus()))
	}
	glog.Info("HandleReceiveMessageAck end.")
}

func HandleNormalMessageAck(conn *net.TCPConn, receivePacket *packet.Packet) {
	glog.Info("HandleNormalMessageAck beginning...")
	normalMessageAck := &common.NormalMessageAck{}
	packet.Unpack(receivePacket, normalMessageAck)
	dao.UpdateNormalMessageStatus(ConnMapUser.Get(conn), normalMessageAck.GetMessageId(), int(normalMessageAck.GetStatus()))
	glog.Info("HandleNormalMessageAck end.")
}

func SendMessage(receiver string, tag uint32, message *common.Message) {
	pac, err := packet.Pack(tag+2, uint32(common.MessageCommand_RECEIVE_MESSAGE), message)
	if err != nil {
		glog.Errorf("Packet: %v", err)
	}

	if dao.UuidCheckOnline(receiver) {
		glog.Infof("%v is online", receiver)
		receiverConn := UserMapConn.Get(receiver).(*net.TCPConn)
		if SendByteStream(receiverConn, pac.GetBytes()) != nil {
			glog.Infof("Add to offline list")
			dao.OfflineMsgAddMsg(receiver, string(pac.GetBytes()))
		}
	} else {
		glog.Infof("%v is offline", receiver)
		dao.OfflineMsgAddMsg(receiver, string(pac.GetBytes()))

		userTokens, err := dao.FindTokensByUserId(receiver)
		glog.Infof("userTokens: %v, %v", userTokens, err)
		pushMessages := []*push.Message{}
		userToken := &push.Message{}
		for _, userToken = range userTokens {
			glog.Infof("User token: %v", userToken)
			pushMessages = append(pushMessages, &push.Message{Token: userToken.Token, Text: "Apns push"})
		}
		push.Send(pushMessages)
	}
}
