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
	uid,_:=c.Get("uid")
	var showTaskservice services.ShowTaskService
	if err:=c.ShouldBindUri(&showTaskservice); err != nil {
		c.JSON(200,e.ErrorReponse{
			Status:e.InvalidParams,
			Msg:e.GetMsg(e.InvalidParams),
			Error: err.Error(),
		})
		return
	}
	res:=showTaskservice.ShowTask(cast.ToUint(uid))
	c.JSON(200,res)
}
func DeleteTask(c *gin.Context)  {
	uid,_:=c.Get("uid")
	var deleteTaskservice services.DeleteTaskService
	if err:=c.ShouldBindUri(&deleteTaskservice); err != nil {
		c.JSON(200,e.ErrorReponse{
			Status:e.InvalidParams,
			Msg:e.GetMsg(e.InvalidParams),
			Error: err.Error(),
		})
		return
	}
	res:=deleteTaskservice.DeleteTask(cast.ToUint(uid))
	c.JSON(200,res)

}
func UpdateTask(c *gin.Context)  {
	tid:=c.Param("tid")
	var updateTaskService services.UpdateTaskService
	if err:=c.ShouldBind(&updateTaskService); err != nil {
		c.JSON(200,e.ErrorReponse{
			Status: e.InvalidParams,
			Msg: e.GetMsg(e.InvalidParams),
			Error: err.Error(),
		})
		return
	}
	res:=updateTaskService.UpdateTask(tid)
	c.JSON(200,res)
}