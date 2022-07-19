package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type SignUp struct {
	User string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	RePassword string `form:"repassword" json:"re_password" binding:"required,eqfield=Password"`
	Email string `form:"email" json:"email" binding:"email"`
	Phone string `form:"phone" json:"phone" binding:"required,min=11,max=11"`
}

func SignUpHandler(c *gin.Context)  {
	var signup SignUp
	log.Println(c.Request.Header)
	err:=c.ShouldBind(&signup)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"msg":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,signup)
}

func CostLogger(c *gin.Context)  {
	start:=time.Now()
	c.Next()
	log.Printf("the request spend time %d ms",time.Since(start).Milliseconds())
}

func CheckToken(c *gin.Context)  {
	token:=c.GetHeader("X-Auth-Token")
	if token!="this is secret"{
		c.JSON(http.StatusOK,gin.H{
			"msg":"Auth failed!",
		})
		c.Abort()
	}
	c.Next()
	{}
}