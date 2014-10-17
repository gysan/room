package db

import (
	"github.com/gysan/room/config"
	"github.com/golang/glog"
	"database/sql"
)

var db *sql.DB

func GetDb() *sql.DB {
	var err error
	if db == nil {
		db, err = sql.Open("mysql", config.MysqlIMSource)
		if err != nil {
			glog.Errorf("sql.Open(\"mysql\", %s) failed (%v)", config.MysqlIMSource, err)
			panic(err)
		}
		glog.Infof("sql.Open(\"mysql\", %s)", config.MysqlIMSource)
	}
	return db
}
