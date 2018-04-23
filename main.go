package main

import (
	"gin-server/api"
	_ "gin-server/config"
	"gin-server/db"
	"gin-server/glog"

	"github.com/gin-gonic/gin"
)

func main() {
	glog.InitLogger()
	db.InitDBConn()
	router := gin.Default()
	router.POST("/", api.HTTPRoute)
	router.Run(":8080")
}
