package main

import (
	"github.com/gin-gonic/gin"
	"github.com/im6/vp3/config"
	"github.com/im6/vp3/handler"
	"github.com/im6/vp3/store"
)

func init() {
	config.LoadConfiguration()
}

func main() {
	r := gin.Default()
  handler.Initialize(r)
	err := store.InitMySQLConn(config.GetSqlConn())
	if err != nil {
		panic(err)
	}

	r.Run(":" + config.GetPort())
}
