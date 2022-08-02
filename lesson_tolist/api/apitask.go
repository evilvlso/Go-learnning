package api

import (
	e "example.com/to_list/pkg/exceptions"
	"example.com/to_list/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func PingService(c *gin.Context) {
	s:=sessions.Default(c)
	s.Set("hello","vdail")
	s.Save()
	c.String(200,"cookie set ok")
}

func ListTask(c *gin.Context)  {
	uid:=c.GetUint("uid")
	var listTaskService services.ListTaskService
	if err:=c.ShouldBindQuery(&listTaskService);err!=nil{
		c.JSON(200,e.ErrorReponse{
			Status:e.InvalidParams,
			Msg:e.GetMsg(e.InvalidParams),
			Error: err.Error(),
		})
		return
	}
	res:=listTaskService.ListTask(uint8(uid))
	c.JSON(200,res)
}

func CreateTask(c *gin.Context)  {
	uid,_:=c.Get("uid")
	var createTaskservice services.CreateTaskService
	if err:=c.ShouldBind(&createTaskservice); err != nil {
		c.JSON(200,e.ErrorReponse{
			Status:e.InvalidParams,
			Msg:e.GetMsg(e.InvalidParams),
			Error: err.Error(),
		})
		return
	}
	res:=createTaskservice.CreateTask(cast.ToUint(uid))
	c.JSON(200,res)
}
func ShowTask(c *gin.Context)  {
	//uid,_:=c.Get("uid")
	//tid:=c.Param("tid")
}
func DeleteTask(c *gin.Context)  {
	//uid,_:=c.Get("uid")
	//tid:=c.Param("tid")

}
func UpdateTask(c *gin.Context)  {
	//uid,_:=c.Get("uid")
	//tid:=c.Param("tid")

}