package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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