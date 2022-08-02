package main

import (
	"example.com/to_list/conf"
	"example.com/to_list/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	quit:=make(chan os.Signal)
	conf.InitMysql()
	router:=router.NewRouter()
	go func() {
		router.Run("127.0.0.1:8080")
	}()
	signal.Notify(quit,syscall.SIGINT)
	<-quit
}
