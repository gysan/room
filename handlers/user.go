package handlers

import (
	"github.com/gysan/room/common"
	"github.com/golang/glog"
	"github.com/gysan/room/dao"
	"net"
	"code.google.com/p/goprotobuf/proto"
	"github.com/gysan/room/packet"
)

func HandleUserLogin(conn *net.TCPConn, receivePacket *packet.Packet) {
	glog.Info("HandleUserLogin beginning...")
	userLogin := &common.UserLogin{}
	packet.Unpack(receivePacket, userLogin)
	userId := userLogin.GetUserId()

	// 检测userId合法性
	if dao.UuidCheckExist(userId) {
		// 如果已经在线则关闭以前的conn
		if dao.UuidCheckOnline(userId) {
			co := UserMapConn.Get(userId)
			if co != nil {
				CloseConn(co.(*net.TCPConn))
			}
		}
		// 上线conn
		InitConn(conn, userId)
	} else {
		CloseConn(conn)
	}

	glog.Infof("UserLogin: %v string: %s", receivePacket, userLogin.String())
	glog.Info("HandleUserLogin end.")

	glog.Info("UserLoginResponse...")
	userLoginResponse := &common.UserLoginResponse{
		Status: proto.Bool(true),
	}
	glog.Infof("UserLoginResponse: %v string: %s", userLoginResponse, userLoginResponse.String())
	err := SendPbData(conn, receivePacket.Tag+1, uint32(common.MessageCommand_USER_LOGIN_RESPONSE), userLoginResponse)
	if err != nil {
		glog.Errorf("Send data error: %v", err)
		return
	}
	glog.Info("UserLoginResponse end.")

	glog.Info("Offline message sending...")
	messages, err := dao.FindOfflineMessages(userId)
	glog.Infof("%v, %v", messages, err)

	v := &common.Message{}
	for _, v = range messages {
		pac, err := packet.Pack(receivePacket.Tag+1, uint32(common.MessageCommand_RECEIVE_MESSAGE), v)
		if err != nil {
			glog.Errorf("Packet: %v", err)
			return
		}
		SendByteStream(conn, pac.GetBytes())
	}
	glog.Info("Offline message sent.")

	glog.Info("Offline normal message sending...")
	normalMessages, err := dao.FindOfflineNormalMessages(userId)
	glog.Infof("%v, %v", normalMessages, err)

	normalMessage := &common.NormalMessage{}
	for _, normalMessage = range normalMessages {
		pac, err := packet.Pack(receivePacket.Tag+1, uint32(common.MessageCommand_NORMARL_MESSAGE), normalMessage)
		if err != nil {
			glog.Errorf("Packet: %v", err)
			return
		}
		SendByteStream(conn, pac.GetBytes())
	}
	glog.Info("Offline normal message sent.")

	token := userLogin.GetToken()
	if token != "" {
		glog.Infof("User Id: [%v] token [%v]", userId, token)
		dao.InsertToken(userId, token)
	}
}

func HandleUserLogout(conn *net.TCPConn, receivePacket *packet.Packet) {
	glog.Info("UserLogoutResponse...")
	userLogoutResponse := &common.UserLogoutResponse{
		Status: proto.Bool(true),
	}
	glog.Infof("UserLogoutResponse: %v string: %s", userLogoutResponse, userLogoutResponse.String())
	err := SendPbData(conn, receivePacket.Tag+1, uint32(common.MessageCommand_USER_LOGOUT_RESPONSE), userLogoutResponse)
	if err != nil {
		glog.Errorf("Send data error: %v", err)
		return
	}
	glog.Info("UserLogoutResponse end.")

	glog.Info("HandleUserLogout beginning...")
	userLogout := &common.UserLogout{}
	packet.Unpack(receivePacket, userLogout)
	CloseConn(conn)
	glog.Info("HandleUserLogout end.")
}
