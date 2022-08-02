package model

import "example.com/to_list/record"

func migrate()  {
	err:=Db.Set("gorm:table_options","charset=utf8mb4").AutoMigrate(&User{},&Task{})
	if err != nil {
		record.Logger.Panicf("table migrate failed err:%s",err.Error())
	}
}
