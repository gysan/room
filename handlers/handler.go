package handlers

import (
	proto "code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/gysan/room/config"
	"github.com/gysan/room/dao"
	"github.com/gysan/room/packet"
	"github.com/gysan/room/utils/safemap"
	"net"
	"time"
	"github.com/golang/glog"
	"github.com/gysan/room/common"
	"strconv"
)

var (
	// 映射conn到user, user --> conn
	UserMapConn *safemap.SafeMap = safemap.NewSafeMap()
	// 映射user到conn, conn --> user
	ConnMapUser *safemap.SafeMap = safemap.NewSafeMap()
	// ConnMapLoginStatus 映射conn的登陆状态,conn->loginstatus
	// loginstatus为nil表示conn已经登陆, 否则 loginstatus表示conn连接服务器时的时间
	// 用于判断登陆是否超时，控制恶意连接
	ConnMapLoginStatus *safemap.SafeMap = safemap.NewSafeMap()
)

//关闭conn
func CloseConn(conn *net.TCPConn) {
	glog.Info("Close connection...")
	conn.Close()
	ConnMapLoginStatus.Delete(conn)
	glog.Infof("Conn login map size: %v, Values: %v", ConnMapLoginStatus.Size(), ConnMapLoginStatus.Items())
	userId := ConnMapUser.Get(conn)
	UserMapConn.Delete(userId)
	glog.Infof("Device conn map size: %v, Values: %v", UserMapConn.Size(), UserMapConn.Items())
	ConnMapUser.Delete(conn)
	glog.Infof("Conn device map size: %v, Values: %v", ConnMapUser.Size(), ConnMapUser.Items())
	if userId != nil {
		dao.UuidOffLine(userId.(string))
		userIdInt, _ := strconv.Atoi(userId.(string))
		dao.UpdateOnline(userIdInt, 0)
	}
	glog.Info("Close connection.")
}

//初始化conn
func InitConn(conn *net.TCPConn, userId string) {
	glog.Info("Init connection...")
	ConnMapLoginStatus.Set(conn, nil)
	glog.Infof("Conn login map size: %v, Values: %v", ConnMapLoginStatus.Size(), ConnMapLoginStatus.Items())
	UserMapConn.Set(userId, conn)
	glog.Infof("Device conn map size: %v, Values: %v", UserMapConn.Size(), UserMapConn.Items())
	ConnMapUser.Set(conn, userId)
	glog.Infof("Conn device map size: %v, Values: %v", ConnMapUser.Size(), ConnMapUser.Items())
	userIdInt, _ := strconv.Atoi(userId)
	dao.UuidOnLine(userId)
	dao.UpdateOnline(userIdInt, 1)
	glog.Info("Init connection.")
}

// 发送字节流
func SendByteStream(conn *net.TCPConn, buf []byte) error {
	conn.SetWriteDeadline(time.Now().Add(time.Duration(config.WriteTimeout) * time.Second))
	n, err := conn.Write(buf)
	if n != len(buf) || err != nil {
		return fmt.Errorf("Write to %v failed, Error: %v", ConnMapUser.Get(conn).(string), err)
	}
	return nil
}

// 发送protobuf结构数据
func SendPbData(conn *net.TCPConn, tag uint32, dataType uint32, pb interface{}) error {
	pac, err := packet.Pack(tag, dataType, pb)
	if err != nil {
		return err
	}
	return SendByteStream(conn, pac.GetBytes())
}

/**
 * 心跳初始化　
 */
func HandleHeartbeatInit(conn *net.TCPConn, receivePacket *packet.Packet) {
	glog.Info("HandleHeartbeatInit beginning...")
	heartbeatInit := &common.HeartbeatInit{}
	packet.Unpack(receivePacket, heartbeatInit)
	glog.Infof("HeartbeatInit: %v string: %s", receivePacket, heartbeatInit.String())
	glog.Info("HandleHeartbeatInit end.")

	glog.Info("HeartbeatInitResponse...")
	heartbeatInitResponse := &common.HeartbeatInitResponse{
		NextHeartbeat: proto.Int32(30),
	}
	glog.Infof("HeartbeatInitResponse: %v string: %s", heartbeatInitResponse, heartbeatInitResponse.String())
	err := SendPbData(conn, receivePacket.Tag+1, uint32(common.MessageCommand_HEARTBEAT_INIT_RESPONSE), heartbeatInitResponse)
	if err != nil {
		glog.Errorf("Send data error: %v", err)
		return
	}
	glog.Info("HeartbeatInitResponse")
}

/**
 * 心跳
 */
func HandleHeartbeat(conn *net.TCPConn, receivePacket *packet.Packet) {
	glog.Info("HandleHeartbeat beginning...")
	heartbeat := &common.Heartbeat{}
	packet.Unpack(receivePacket, heartbeat)
	glog.Infof("Heartbeat: %v string: %s", receivePacket, heartbeat.String())
	glog.Info("HandleHeartbeat end.")

	conn.SetReadDeadline(time.Now().Add(time.Duration(heartbeat.GetLastDelay() + 5) * time.Second))

	glog.Info("HeartbeatResponse...")
	heartbeatResponse := &common.HeartbeatResponse{
		NextHeartbeat: proto.Int32(heartbeat.GetLastDelay() + 5),
	}
	glog.Infof("HeartbeatResponse: %v string: %s", heartbeatResponse, heartbeatResponse.String())
	err := SendPbData(conn, receivePacket.Tag+1, uint32(common.MessageCommand_HEARTBEAT_RESPONSE), heartbeatResponse)
	if err != nil {
		glog.Errorf("Send data error: %v", err)
		return
	}
	glog.Info("HeartbeatResponse end.")
}
