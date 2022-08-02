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
	err:=model.Db.Select("id,title,comment,status").Where("uid=?",uid).Find(&l.Tasks).
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
	}
}

type ShowTaskService struct {
}

type DeleteTaskService struct {
}

type UpdateTaskService struct {
}
