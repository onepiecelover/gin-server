package api

import (
	"encoding/json"
	"gin-server/glog"
	"log"
	"net/http"
	"uframework/common"

	"github.com/gin-gonic/gin"
)

type EReq struct {
	Name string
	Age  int
}

func Echo(c *gin.Context, data interface{}) {
	log.Println("in Echo", ufcommon.GetGID())
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Println("fuck Marshal")
		c.JSON(http.StatusUnauthorized, gin.H{"status": "action is not found"})
		return
	}
	var req EReq
	err = json.Unmarshal(jsonBytes, &req)
	if err != nil {
		log.Println("fuck unMarshal")
		c.JSON(http.StatusUnauthorized, gin.H{"status": "action is not found"})
		return
	}
	log.Println(req)
	glog.DEBUG(req)
	glog.INFOF("hehehe %s", req.Name)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	//time.Sleep(10 * time.Second)
	return
}
