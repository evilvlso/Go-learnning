package api

import (
	e "example.com/to_list/pkg/exceptions"
	"example.com/to_list/record"
	"example.com/to_list/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginService services.LoginService
	if err:=c.ShouldBind(&loginService);err!=nil{
		c.JSON(200,e.ErrorReponse{
			Status:e.InvalidParams,
			Msg:e.GetMsg(e.InvalidParams),
			Error: err.Error(),
		})
		return
	}
	res:=loginService.Login()
	c.JSON(200,res)
}

func SignUp(c *gin.Context) {
	var signUpService services.SignUpService
	if err:=c.ShouldBind(&signUpService);err!=nil{
		record.Logger.Printf("SignUp validate failed!")
		c.JSON(200,e.ErrorReponse{
			Status:e.InvalidParams,
			Msg:e.GetMsg(e.InvalidParams),
			Error: err.Error(),
		})
		return
	}
	res:=signUpService.Register()
	c.JSON(200,res)
}
