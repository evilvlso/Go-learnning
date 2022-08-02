package conf

import (
	"example.com/to_list/model"
	"example.com/to_list/record"
	"fmt"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/go-ini/ini"
)

func InitMysql() error {
	cfg,err:=ini.Load("conf/conf.ini")
	if err != nil {
		record.Logger.Panicf("Config loading error:%s",err.Error())
	}
	mysqlSection:=cfg.Section("mysql")
	dsn:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		mysqlSection.Key("DbUser").String(),
		mysqlSection.Key("DbPasswd").String(),
		mysqlSection.Key("DbHost").String(),
		mysqlSection.Key("DbPort").String(),
		mysqlSection.Key("Db").String())
	model.InitDb(dsn)
	return nil
}

func NewStore() cookie.Store {
	cfg,err:=ini.Load("conf/conf.ini")
	if err != nil {
		record.Logger.Printf("NewStore Config loading error:%s",err.Error())
	}
	redisSection:=cfg.Section("redis")
	store,err:=redis.NewStore(redisSection.Key("idleSize").MustInt(),
		redisSection.Key("Network").String(),
		redisSection.Key("Address").String(),
		redisSection.Key("Passwd").String(),
		[]byte(redisSection.Key("Secret").String()))
	if err != nil {
		record.Logger.Panicf("Redis Store Init error:%s",err.Error())
		store=cookie.NewStore([]byte("vdail"))
	}
	return store
}