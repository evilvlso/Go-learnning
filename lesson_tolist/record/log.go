package record

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Logger *log.Logger
)

func init()  {
	fmt.Println("initialize logger")
	writer:=os.Stdout
	writer2,_:= os.OpenFile("./logs/log.log",os.O_APPEND|os.O_CREATE,0755)
	defer writer2.Close()
	Logger=log.New(io.MultiWriter(writer2,writer),"",os.O_APPEND|os.O_CREATE)
}