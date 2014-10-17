// Code generated by protoc-gen-go.
// source: MessageRouting.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	MessageRouting.proto

It has these top-level messages:
	HeartbeatInit
	HeartbeatInitResponse
	Heartbeat
	HeartbeatResponse
	UserLogin
	UserLoginResponse
	UserLogout
	UserLogoutResponse
	Message
	MessageResponse
	ReceiveMessageAck
	NormalMessage
	NormalMessageAck
*/
package common

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// 消息类型
type MessageCommand int32

const (
	// 心跳初始化
	MessageCommand_HEARTBEAT_INIT MessageCommand = 0
	// 心跳
	MessageCommand_HEARTBEAT MessageCommand = 1
	// 用户登录
	MessageCommand_USER_LOGIN MessageCommand = 2
	// 用户登出
	MessageCommand_USER_LOGOUT MessageCommand = 3
	// 消息
	MessageCommand_MESSAGE MessageCommand = 4
	// 接收到消息后向服务器发确认
	MessageCommand_RECEIVE_MESSAGE_ACK MessageCommand = 5
	// Normal消息确认
	MessageCommand_NORMARL_MESSAGE_ACK MessageCommand = 6
	// 心跳初始化响应
	MessageCommand_HEARTBEAT_INIT_RESPONSE MessageCommand = 100
	// 心跳响应
	MessageCommand_HEARTBEAT_RESPONSE MessageCommand = 101
	// 用户登录响应
	MessageCommand_USER_LOGIN_RESPONSE MessageCommand = 102
	// 用户登出响应
	MessageCommand_USER_LOGOUT_RESPONSE MessageCommand = 103
	// 消息响应
	MessageCommand_MESSAGE_RESPONSE MessageCommand = 104
	// 接收服务器发的消息
	MessageCommand_RECEIVE_MESSAGE MessageCommand = 200
	// Normarl push
	MessageCommand_NORMARL_MESSAGE MessageCommand = 201
)

var MessageCommand_name = map[int32]string{
	0:   "HEARTBEAT_INIT",
	1:   "HEARTBEAT",
	2:   "USER_LOGIN",
	3:   "USER_LOGOUT",
	4:   "MESSAGE",
	5:   "RECEIVE_MESSAGE_ACK",
	6:   "NORMARL_MESSAGE_ACK",
	100: "HEARTBEAT_INIT_RESPONSE",
	101: "HEARTBEAT_RESPONSE",
	102: "USER_LOGIN_RESPONSE",
	103: "USER_LOGOUT_RESPONSE",
	104: "MESSAGE_RESPONSE",
	200: "RECEIVE_MESSAGE",
	201: "NORMARL_MESSAGE",
}
var MessageCommand_value = map[string]int32{
	"HEARTBEAT_INIT":          0,
	"HEARTBEAT":               1,
	"USER_LOGIN":              2,
	"USER_LOGOUT":             3,
	"MESSAGE":                 4,
	"RECEIVE_MESSAGE_ACK":     5,
	"NORMARL_MESSAGE_ACK":     6,
	"HEARTBEAT_INIT_RESPONSE": 100,
	"HEARTBEAT_RESPONSE":      101,
	"USER_LOGIN_RESPONSE":     102,
	"USER_LOGOUT_RESPONSE":    103,
	"MESSAGE_RESPONSE":        104,
	"RECEIVE_MESSAGE":         200,
	"NORMARL_MESSAGE":         201,
}

func (x MessageCommand) Enum() *MessageCommand {
	p := new(MessageCommand)
	*p = x
	return p
}
func (x MessageCommand) String() string {
	return proto.EnumName(MessageCommand_name, int32(x))
}
func (x *MessageCommand) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageCommand_value, data, "MessageCommand")
	if err != nil {
		return err
	}
	*x = MessageCommand(value)
	return nil
}

// 0: 心跳初始化
type HeartbeatInit struct {
	// 上次长连接最后一次心跳的间隔（即便不成功也算数)
	LastTimeout *int32 `protobuf:"varint,1,req,name=last_timeout" json:"last_timeout,omitempty"`
	// 扩展字段
	Extra            *string `protobuf:"bytes,2,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HeartbeatInit) Reset()         { *m = HeartbeatInit{} }
func (m *HeartbeatInit) String() string { return proto.CompactTextString(m) }
func (*HeartbeatInit) ProtoMessage()    {}

func (m *HeartbeatInit) GetLastTimeout() int32 {
	if m != nil && m.LastTimeout != nil {
		return *m.LastTimeout
	}
	return 0
}

func (m *HeartbeatInit) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

type HeartbeatInitResponse struct {
	// 返回下一次心跳时间
	NextHeartbeat *int32 `protobuf:"varint,1,opt,name=next_heartbeat" json:"next_heartbeat,omitempty"`
	// 扩展字段
	Extra            *string `protobuf:"bytes,2,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HeartbeatInitResponse) Reset()         { *m = HeartbeatInitResponse{} }
func (m *HeartbeatInitResponse) String() string { return proto.CompactTextString(m) }
func (*HeartbeatInitResponse) ProtoMessage()    {}

func (m *HeartbeatInitResponse) GetNextHeartbeat() int32 {
	if m != nil && m.NextHeartbeat != nil {
		return *m.NextHeartbeat
	}
	return 0
}

func (m *HeartbeatInitResponse) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

// 1: 心跳
type Heartbeat struct {
	// 　上次心跳时间
	LastDelay *int32 `protobuf:"varint,1,opt,name=last_delay" json:"last_delay,omitempty"`
	// 　扩展字段
	Extra            *string `protobuf:"bytes,2,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Heartbeat) Reset()         { *m = Heartbeat{} }
func (m *Heartbeat) String() string { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()    {}

func (m *Heartbeat) GetLastDelay() int32 {
	if m != nil && m.LastDelay != nil {
		return *m.LastDelay
	}
	return 0
}

func (m *Heartbeat) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

// 101: 心跳响应
type HeartbeatResponse struct {
	// 返回下一次心跳时间
	NextHeartbeat *int32 `protobuf:"varint,1,opt,name=next_heartbeat" json:"next_heartbeat,omitempty"`
	// 扩展字段
	Extra            *string `protobuf:"bytes,2,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HeartbeatResponse) Reset()         { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string { return proto.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()    {}

func (m *HeartbeatResponse) GetNextHeartbeat() int32 {
	if m != nil && m.NextHeartbeat != nil {
		return *m.NextHeartbeat
	}
	return 0
}

func (m *HeartbeatResponse) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

// 2: 用户登录
type UserLogin struct {
	// 用户id
	UserId *string `protobuf:"bytes,1,req,name=user_id" json:"user_id,omitempty"`
	// 应用id
	AppId *int32 `protobuf:"varint,2,opt,name=app_id" json:"app_id,omitempty"`
	// Channel
	Channel *string `protobuf:"bytes,3,opt,name=channel" json:"channel,omitempty"`
	// 版本
	Version *string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	// token
	Token *string `protobuf:"bytes,5,opt,name=token" json:"token,omitempty"`
	// 扩展字段
	Extra            *string `protobuf:"bytes,6,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserLogin) Reset()         { *m = UserLogin{} }
func (m *UserLogin) String() string { return proto.CompactTextString(m) }
func (*UserLogin) ProtoMessage()    {}

func (m *UserLogin) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *UserLogin) GetAppId() int32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *UserLogin) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *UserLogin) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *UserLogin) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *UserLogin) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

// 102: 用户登录响应
type UserLoginResponse struct {
	// 用户登录是否成功
	Status *bool `protobuf:"varint,1,req,name=status" json:"status,omitempty"`
	// 扩展字段
	Extra            *string `protobuf:"bytes,2,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserLoginResponse) Reset()         { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()    {}

func (m *UserLoginResponse) GetStatus() bool {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return false
}

func (m *UserLoginResponse) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

// 3: 用户登出
type UserLogout struct {
	// 用户id
	UserId *string `protobuf:"bytes,1,req,name=user_id" json:"user_id,omitempty"`
	// 扩展字段
	Extra            *string `protobuf:"bytes,2,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserLogout) Reset()         { *m = UserLogout{} }
func (m *UserLogout) String() string { return proto.CompactTextString(m) }
func (*UserLogout) ProtoMessage()    {}

func (m *UserLogout) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *UserLogout) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

// 103: 用户登出返回
type UserLogoutResponse struct {
	// 用户是否成功退出
	Status *bool `protobuf:"varint,1,req,name=status" json:"status,omitempty"`
	// 扩展字段
	Extra            *string `protobuf:"bytes,2,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserLogoutResponse) Reset()         { *m = UserLogoutResponse{} }
func (m *UserLogoutResponse) String() string { return proto.CompactTextString(m) }
func (*UserLogoutResponse) ProtoMessage()    {}

func (m *UserLogoutResponse) GetStatus() bool {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return false
}

func (m *UserLogoutResponse) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

// 4, 200: 消息, 消息响应
type Message struct {
	// 客户端生成消息id
	MessageId *string `protobuf:"bytes,1,req,name=message_id" json:"message_id,omitempty"`
	// 发送者
	Sender *string `protobuf:"bytes,2,req,name=sender" json:"sender,omitempty"`
	// 接收者
	Receiver *string `protobuf:"bytes,3,req,name=receiver" json:"receiver,omitempty"`
	// 消息类型
	Type *string `protobuf:"bytes,4,req,name=type" json:"type,omitempty"`
	// 消息体
	Body *string `protobuf:"bytes,5,req,name=body" json:"body,omitempty"`
	// 收发消息时间
	Date *int64 `protobuf:"varint,6,opt,name=date" json:"date,omitempty"`
	// 扩展字段
	Extra            *string `protobuf:"bytes,7,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

func (m *Message) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *Message) GetSender() string {
	if m != nil && m.Sender != nil {
		return *m.Sender
	}
	return ""
}

func (m *Message) GetReceiver() string {
	if m != nil && m.Receiver != nil {
		return *m.Receiver
	}
	return ""
}

func (m *Message) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Message) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *Message) GetDate() int64 {
	if m != nil && m.Date != nil {
		return *m.Date
	}
	return 0
}

func (m *Message) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

// 104: 消息响应
type MessageResponse struct {
	// 客户端生产消息id
	MessageId *string `protobuf:"bytes,1,req,name=message_id" json:"message_id,omitempty"`
	// 状态
	Status *bool `protobuf:"varint,2,req,name=status" json:"status,omitempty"`
	// 扩展字段
	Extra            *string `protobuf:"bytes,3,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}

func (m *MessageResponse) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *MessageResponse) GetStatus() bool {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return false
}

func (m *MessageResponse) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

// 5: 接收到消息后向服务器发确认
type ReceiveMessageAck struct {
	// 消息id
	MessageId *string `protobuf:"bytes,1,req,name=message_id" json:"message_id,omitempty"`
	// 消息类型
	Type *string `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
	// 消息状态　0：正在下发中；1：已经下发push但未查看；2：收到push并已经查看
	Status *int32 `protobuf:"varint,3,req,name=status" json:"status,omitempty"`
	// 扩展字段
	Extra            *string `protobuf:"bytes,4,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReceiveMessageAck) Reset()         { *m = ReceiveMessageAck{} }
func (m *ReceiveMessageAck) String() string { return proto.CompactTextString(m) }
func (*ReceiveMessageAck) ProtoMessage()    {}

func (m *ReceiveMessageAck) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *ReceiveMessageAck) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *ReceiveMessageAck) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *ReceiveMessageAck) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

// 201: 通用PUSH
type NormalMessage struct {
	// 消息id
	MessageId *string `protobuf:"bytes,1,req,name=message_id" json:"message_id,omitempty"`
	// 收件人
	Receiver *string `protobuf:"bytes,2,req,name=receiver" json:"receiver,omitempty"`
	// 内容， 最长1024字节
	Content []byte `protobuf:"bytes,3,req,name=content" json:"content,omitempty"`
	// 消息生成的时间和日期， 默认为东八区
	Date *int64 `protobuf:"varint,4,req,name=date" json:"date,omitempty"`
	// 过期时间,单位：s
	Expire *int32 `protobuf:"varint,5,opt,name=expire" json:"expire,omitempty"`
	// 客户端收到消息确认
	Extra            *string `protobuf:"bytes,6,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NormalMessage) Reset()         { *m = NormalMessage{} }
func (m *NormalMessage) String() string { return proto.CompactTextString(m) }
func (*NormalMessage) ProtoMessage()    {}

func (m *NormalMessage) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *NormalMessage) GetReceiver() string {
	if m != nil && m.Receiver != nil {
		return *m.Receiver
	}
	return ""
}

func (m *NormalMessage) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *NormalMessage) GetDate() int64 {
	if m != nil && m.Date != nil {
		return *m.Date
	}
	return 0
}

func (m *NormalMessage) GetExpire() int32 {
	if m != nil && m.Expire != nil {
		return *m.Expire
	}
	return 0
}

func (m *NormalMessage) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

// 6: 单条NORMAL消息ack
type NormalMessageAck struct {
	// 消息id
	MessageId *string `protobuf:"bytes,1,req,name=message_id" json:"message_id,omitempty"`
	// 消息状态　0：正在下发中；1：已经下发push但未查看；2：收到push并已经查看
	Status *int32 `protobuf:"varint,2,req,name=status" json:"status,omitempty"`
	// 扩展字段
	Extra            *string `protobuf:"bytes,3,opt,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NormalMessageAck) Reset()         { *m = NormalMessageAck{} }
func (m *NormalMessageAck) String() string { return proto.CompactTextString(m) }
func (*NormalMessageAck) ProtoMessage()    {}

func (m *NormalMessageAck) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *NormalMessageAck) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *NormalMessageAck) GetExtra() string {
	if m != nil && m.Extra != nil {
		return *m.Extra
	}
	return ""
}

func init() {
	proto.RegisterEnum("common.MessageCommand", MessageCommand_name, MessageCommand_value)
}