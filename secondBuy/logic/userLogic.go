package logic

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type User struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}


var db = initDb()

func initDb() *sql.DB {
	db, err := sql.Open("mysql", "root:wananle@(127.0.0.1:3306)/second_buy")
	checkErr(err, "sql.Open failed")

	return db
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func checkParam(c *gin.Context) bool{

	var user User
	if c.ShouldBind(&user) == nil {
		return true
	}else {
		return false
	}
}


func Select(c *gin.Context) []User{

	users := make([]User, 0)

	if !checkParam(c) {
		log.Fatalln("参数校验错误")
	}

	userName := c.PostForm("userName")
	res, err := db.Query("SELECT user_name,account,password FROM user where user_name=?",userName)

	defer res.Close()
	checkErr(err, "select user error")
	for res.Next() {
		var user User
		res.Scan(&user.UserName, &user.Account, &user.Password)
		fmt.Println(user)
		users = append(users, user)
	}

	if err = res.Err(); err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Users": users,
	})

	return users
}

func Insert(c *gin.Context,param []string) int64{

	users := make([]User, 0)
	fmt.Println(users)

	for key,value :=range param {
		fmt.Println(key)
		fmt.Println(value)
	}

	userName := c.PostForm("user_name")
	password := c.PostForm("password")
	if userName == "" {
		fmt.Println("username is empty")
		os.Exit(1)
	}
	if password == "" {
		fmt.Println("password is empty")
		os.Exit(1)
	}
	fmt.Println(userName)
	fmt.Println(password)
	res, err := db.Exec("insert into user (user_name,account,password)values (?,?,?)",userName,0,password)
	id, err := res.LastInsertId()
	checkErr(err, "select user error")
	fmt.Println("insert Id is %d",id)

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
	return id
}
