package config

import (
	"fmt"
	"github.com/robfig/config"
	"github.com/golang/glog"
)

var (
	Addr                string
	Http                string
	AcceptTimeout       int
	ReadTimeout         int
	WriteTimeout        int

	MysqlIMSource       string

	ApnsUrl             string
	ApnsCertificate     string
	ApnsKey             string

	RedisBind           string

	BelieveRoomUserLimit  int

	UuidDB              string
	OfflineMsgidsDB     string
	IdToMsgDB           string
	TimedUpdateDB       int
	LogFile             string
)

func ReadIniFile(iniFile string) error {
	conf, err := config.ReadDefault(iniFile)
	if err != nil {
		return fmt.Errorf("Read %v error. %v", iniFile, err.Error())
	}

	Addr, _ = conf.String("service", "addr")
	Http, _ = conf.String("service", "http")
	AcceptTimeout, _ = conf.Int("service", "accept_timeout")
	ReadTimeout, _ = conf.Int("service", "read_timeout")
	WriteTimeout, _ = conf.Int("service", "write_timeout")

	MysqlIMSource, _ = conf.String("mysql", "mysql_im_source")

	ApnsUrl, _ = conf.String("apns", "apns_url")
	ApnsCertificate, _ = conf.String("apns", "apns_certificate")
	ApnsKey, _ = conf.String("apns", "apns_key")

	RedisBind, _ = conf.String("redis", "redis_bind")

	BelieveRoomUserLimit, _ = conf.Int("believe", "believe_room_user_limit")
	glog.Infof("-0-0-0-0-0-0:%v", BelieveRoomUserLimit)
	UuidDB, _ = conf.String("data", "uuid_db")
	OfflineMsgidsDB, _ = conf.String("data", "offline_msgids_db")
	IdToMsgDB, _ = conf.String("data", "id_to_msg_db")
	TimedUpdateDB, _ = conf.Int("data", "timed_update_db")

	LogFile, _ = conf.String("debug", "logfile")

	return nil
}
