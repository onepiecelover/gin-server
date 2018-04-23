package api

import (
	"encoding/json"
	"gin-server/glog"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestJson struct {
	Action string `json:"Action"`
	Data   interface{}
}

var APIList = map[string]struct {
	Action string
	Run    func(c *gin.Context, d interface{})
}{
	"Echo": {"Echo", Echo},
}

func (r *RequestJson) server(c *gin.Context) {
	f, ok := APIList[r.Action]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "action is not found"})
		return
	}
	f.Run(c, r.Data)
}

func HTTPRoute(c *gin.Context) {
	var r RequestJson
	if c.BindJSON(&r) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	log.Println(r.Action)
	log.Println(r.Data)
	glog.INFO(r.Action)
	s, _ := json.MarshalIndent(r.Data, "", "    ")
	glog.DEBUG(string(s))
	r.server(c)
	return
}
