package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getMsg(c *gin.Context) {

	c.String(http.StatusOK, "terry get go")
}

func main() {

	router := gin.Default()

	//RESTful路由
	router.GET("/RESTful", getMsg)

	router.Run("127.0.0.1:8082")

}
