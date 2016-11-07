package main

import (
  "net/http"
  "gopkg.in/gin-gonic/gin.v1"
  )

  type User struct {
    Id int64 `db:"id" json:"id"`
    Firstname string `db:"firstname" json:"firstname"`
    Lastname string `db:"lastname" json:"lastname"`
  }

  func main() {
	 router := gin.Default()
	 router.GET("/index", func(c *gin.Context) {
		type Users []User
	    var users = Users{
           User{Id: 1, Firstname: "Oliver", Lastname: "Queen"},
		   User{Id: 2, Firstname: "Malcom", Lastname: "Merlyn"},
		}
		c.JSON(http.StatusOK, users)
	 })
	 router.Run()
  }
