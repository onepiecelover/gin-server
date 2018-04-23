package db

import (
	"database/sql"
	"gin-server/glog"
	"time"

	manager "github.com/didi/gendry/manager"
	"github.com/spf13/viper"
)

func InitDBConn() *sql.DB {
	dbName := viper.GetString("db.mysql.name")
	user := viper.GetString("db.mysql.user")
	password := viper.GetString("db.mysql.password")
	host := viper.GetString("db.mysql.host")
	db, err := manager.New(dbName, user, password, host).Set(
		manager.SetCharset("utf8"),
		manager.SetAllowCleartextPasswords(true),
		manager.SetInterpolateParams(true),
		manager.SetTimeout(1*time.Second),
		manager.SetReadTimeout(1*time.Second),
	).Port(3306).Open(true)
	if err != nil {
		glog.DEBUG("fuck db")
		return nil
	}
	glog.DEBUG("InitDBConn Success")
	return db
}
