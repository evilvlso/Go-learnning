package services

import (
	"example.com/to_list/model"
	e "example.com/to_list/pkg/exceptions"
	"example.com/to_list/pkg/utils"
	"example.com/to_list/serialize"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

const (
	dumplicateErrorCode = 1062
)
type LoginService struct {
	NickName string `json:"nickname" form:"nick_name" binding:"required,min=4,max=16"`
	Passwd   string	`json:"passwd" form:"passwd" binding:"required"`
}

func (s LoginService) Login() serialize.Response {
	nickName:=s.NickName
	var user model.User
	//检查用户是否存在
	if err:=model.Db.Where("nick_name=?",nickName).First(&user).Error;err != nil {
		if gorm.ErrRecordNotFound==err{
			return serialize.Response{
				Status: e.ErrorNotExistUser,
				Msg: e.GetMsg(e.ErrorNotExistUser),
			}
		}else{
			return serialize.Response{
				Status: e.ERROR,
				Msg: e.GetMsg(e.ERROR),
				Error: err.Error(),
			}
		}
	}
	//校验密码
	passwd:=s.Passwd
	if ok:=user.CheckPasswd(passwd);ok{
		token:=utils.GenerateToken(utils.User{Uid:user.ID})
		data:=serialize.Token{Data:token}
		return serialize.Response{
			Status: e.SUCCESS,
			Msg:e.GetMsg(e.SUCCESS),
			Data: data.SerializeToken(),
		}
	}
	return serialize.Response{
		Status: e.ErrorNotCompare,
		Msg:e.GetMsg(e.ErrorNotCompare),
	}
}

type SignUpService struct {
	NickName string `form:"nick_name" json:"nickname" binding:"required,min=4,max=16"`
	Passwd   string	`json:"passwd" form:"passwd" binding:"required,alphanum"`
	RePasswd string `json:"repasswd" form:"repasswd" binding:"required,eqfield=Passwd,alphanum"`
	Email	string `json:"email" form:"email" binding:"required,email"`
}

func (s *SignUpService) Register() serialize.Response {
	var user model.User
	user.NickName=s.NickName
	user.CryptPasswd(s.Passwd)
	err:=model.Db.Create(&user).Error
	if err != nil && err.(*mysql.MySQLError).Number==dumplicateErrorCode{
		return serialize.Response{
			Status: e.ErrorExistUser,
			Msg: e.GetMsg(e.ErrorExistUser),
		}
	}else if err!=nil{
		return serialize.Response{
			Status: e.ERROR,
			Msg: e.GetMsg(e.ERROR),
			Error: err.Error(),
		}
	}


	//var count int64
	//err:=model.Db.Where("nick_name=?",s.NickName).Find(&user).Count(&count).Error
	//if err != nil {
	//	record.Logger.Printf("数据库查询失败! err:%s",err.Error())
	//	return serialize.Response{
	//		Status: e.ERROR,
	//		Msg: e.GetMsg(e.ERROR),
	//		Error: err.Error(),
	//	}
	//}
	//if count!=0{
	//	record.Logger.Printf("%s user has existed!",s.NickName)
	//	return serialize.Response{
	//		Status: e.ErrorExistUser,
	//		Msg: e.GetMsg(e.ErrorExistUser),
	//	}
	//}
	//user.NickName=s.NickName
	//err=user.CryptPasswd(s.Passwd)
	//model.Db.Create(&user)
	//record.Logger.Printf("%s register sucess",s.NickName)
	return serialize.Response{
		Status: e.SUCCESS,
		Msg: e.GetMsg(e.SUCCESS),
	}

}