package model

import (
	"example.com/to_list/record"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var (
	Db *gorm.DB
)

func InitDb(dsn string)  {
	db,err:=gorm.Open(
		mysql.New(mysql.Config{
			DSN:dsn,
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				NoLowerCase:false,
			},
			Logger:logger.Default.LogMode(logger.Info),
		})
	if err != nil {
		record.Logger.Panic("Mysql connect failed err:%s",err.Error())
	}
	sqlDb,err:=db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(20)
	sqlDb.SetConnMaxLifetime(time.Second*1800)
	Db=db
	migrate()
}