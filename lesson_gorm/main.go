package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code string
	Price uint

}
func main() {
	log.SetFlags(log.Lshortfile|log.Ldate)
	db,err:=gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/douban?parseTime=True&loc=Local"),&gorm.Config{})
	if err != nil {
		log.Panicf("failed to connect db, error:%s",err.Error())
	}
	err=db.AutoMigrate(&Product{})
	if err != nil {
		log.Println(err)
	}
	//db.Create(&Product{Code: "idj23",Price: 300})

	var product Product
	////log.Printf("%T-%+v",&product,&product)
	//db.Find(&product,"id>=?",2)
	//log.Printf("%T-%+v",&product,&product)
	//db.Model(&product).Updates(&Product{Price: 800})
	//db.Find(&product,"id>=?",2)
	//log.Printf("%T-%+v",&product,&product)
	//db.Model(&product).Updates(map[string]interface{}{"Code":"FF105"})
	var count int64
	err=db.Debug().Where("price=?",1000).First(&product).Error
	if err != nil {
		log.Println(err==gorm.ErrRecordNotFound)
	}
	log.Printf("%T-%+v",&product,&product)
	log.Printf("%d",count)
	//db.Delete(&product)


}
