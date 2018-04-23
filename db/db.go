package db

import (
	"database/sql"
	"gin-server/glog"
	"log"

	manager "github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func InitMysqlDBConn() *sql.DB {
	dbName := viper.GetString("db.mysql.name")
	user := viper.GetString("db.mysql.user")
	password := viper.GetString("db.mysql.password")
	host := viper.GetString("db.mysql.host")
	db, err := manager.New(dbName, user, password, host).Set(
		manager.SetCharset("utf8"),
		manager.SetAllowCleartextPasswords(true),
		manager.SetInterpolateParams(true),
		//manager.SetTimeout(1*time.Second),
		//manager.SetReadTimeout(1*time.Second),
	).Port(3306).Open(true)
	if err != nil {
		glog.DEBUG("fuck db", err)
		log.Println("init failed", err)
		return nil
	}
	log.Println("init success")
	glog.DEBUG("InitDBConn Success")
	return db
}
