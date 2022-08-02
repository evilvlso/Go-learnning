package services

import (
	"example.com/to_list/model"
	e "example.com/to_list/pkg/exceptions"
	"example.com/to_list/serialize"
)

type ListTaskService struct {
	Tasks []model.Task `json:"tasks,omitempty" form:"tasks,omitempty"`
	Limit int `json:"limit,omitempty" form:"limit,default=10"`
	Offset int `json:"offset,omitempty" form:"offset,default=0"`

}

func (l ListTaskService) ListTask(uid uint8) serialize.Response {
	err:=model.Db.Select("id,title,comment,status").Where("uid=?",uid).Order("created_at DESC").Find(&l.Tasks).
		Limit(l.Limit).Offset(l.Offset).Error
	if err != nil {
		return serialize.Response{
			Status:e.ErrorDatabase,
			Msg:e.GetMsg(e.ErrorDatabase),
			Error: err.Error(),
		}
	}
	return serialize.Response{
		Status:e.SUCCESS,
		Msg:e.GetMsg(e.SUCCESS),
		Data:serialize.SerializeTask(l.Tasks),
	}
}

type CreateTaskService struct {
	Title string `json:"title" form:"title" binding:"required"`
	Comment string `json:"comment" form:"comment" binding:"required"`
	Status uint8 `json:"status" form:"status" `

}

func (s CreateTaskService) CreateTask(uid uint) serialize.Response {
	newTask:= model.Task{
		Uid: uid,
		Title: s.Title,
		Comment: s.Comment,
		Status: s.Status,
	}
	if err:=model.Db.Create(&newTask).Error;err!=nil{
		return serialize.Response{
			Status:e.ErrorDatabase,
			Msg:e.GetMsg(e.ErrorDatabase),
			Error: err.Error(),
		}
	}
	return serialize.Response{
		Status:e.SUCCESS,
		Msg:e.GetMsg(e.SUCCESS),
		//应该return tid
		Data:serialize.DataList{},
	}
}

type ShowTaskService struct {
	Tid uint `json:"tid" form:"tid" uri:"tid" binding:"required"`
}

func (s ShowTaskService) ShowTask(uid uint) serialize.Response {
	var task model.Task
	if err:=model.Db.Select([]string{"title","comment","status"}).Where("uid=? and id=?",uid,s.Tid).First(&task).Error;err!=nil{
		return serialize.Response{
			Status: e.ErrorDatabase,
			Msg: e.GetMsg(e.ErrorDatabase),
			Error: err.Error(),
		}
	}
	return serialize.Response{
		Status: e.SUCCESS,
		Msg: e.GetMsg(e.SUCCESS),
		Data: serialize.SerializeTask([]model.Task{task}),
	}
}

type DeleteTaskService struct {
	Tid uint `json:"tid" form:"tid"  uri:"tid" binding:"required"`
}

func (s DeleteTaskService) DeleteTask(uid uint) serialize.Response {
	var task model.Task
	if err:=model.Db.Delete(&task,"uid=? and id=?",uid,s.Tid).Error;err!=nil{
		return serialize.Response{
			Status: e.ErrorDatabase,
			Msg: e.GetMsg(e.ErrorDatabase),
			Error: err.Error(),
		}
	}
	return serialize.Response{
		Status: e.SUCCESS,
		Msg: e.GetMsg(e.SUCCESS),
		Data: "删除成功",
	}
}

type UpdateTaskService struct {
	Title string `json:"title" form:"title" binding:"required"`
	Comment string `json:"comment" form:"comment" binding:"required"`
	Status uint8 `json:"status" form:"status" binding:"required"`
	Tid   uint `json:"tid" form:"tid"  uri:"tid"`
}

func (u UpdateTaskService) UpdateTask(tid string) serialize.Response {
	var task model.Task
	if err:=model.Db.Where("id=?",tid).First(&task).Error; err != nil {
		return serialize.Response{
					Status: e.ErrorDatabase,
					Msg: e.GetMsg(e.ErrorDatabase),
					Error: err.Error(),
				}
	}
	task.Title=u.Title
	task.Comment=u.Comment
	task.Status=u.Status
	model.Db.Save(&task)
	return serialize.Response{
		Status: e.SUCCESS,
		Msg: e.GetMsg(e.SUCCESS),
		Data: "更新成功",
	}
}