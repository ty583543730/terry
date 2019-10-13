package main

import (
	"./app"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()
	router.POST("/register", func(c *gin.Context) {

		userName := c.PostForm("user_name")

		fmt.Println(userName)

		c.JSON(200, gin.H{
			"status":   "posted",
			"userName": userName,
		})

		users := app.Select(c)
		if len(users) > 0 {
			fmt.Println("username is already exist")
			os.Exit(1)
		}

		param := []string{"user_name","account","password"}
		res := app.Insert(c,param)

		if res < 0 {
			fmt.Println("create account error")
			os.Exit(1)
		}

	})

	router.Run(":8081")
}
