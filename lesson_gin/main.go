package main

import (
	example "example.com/gin/pb"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"os"
	"os/signal"
	"syscall"

	// "github.com/gogo/protobuf/proto"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
)

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	log.Printf("%T-%v", cast.ToInt8(c.Query("limit")), c.Query("limit"))
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"action": action,
	})
}

type Student struct {
	Sid    int8      `json:"sid"`
	Name   string   `json:"name"`
	Age    int      `json:"age,omitempty"`
	Course []string `json:"course"`
}

func ReturnJson(c *gin.Context) {
	sid := c.Param("sid")
	s := &Student{cast.ToInt8(sid), "Minty", 18, []string{"english", "math", "psycology"}}
	log.Println(s)
	c.JSON(http.StatusOK, s)
}

func ReturnProto(c *gin.Context) {
	et := &example.Teacher{Name:"Bai", Age:40, Level: example.Teacher_MASTER, Course:"english"}
	b,err:=proto.Marshal(et)
	if err != nil {
		log.Panic(err.Error())
	}
	log.Printf("%T,%v",b,b)
	c.ProtoBuf(http.StatusOK,et)
}
func main() {
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//router:=gin.New()
	router.RouterGroup.Handlers=append(router.RouterGroup.Handlers,CostLogger,CheckToken)
	ProductList := router.Group("/product")
	ProductList.GET("/:id/:action", GetProduct)

	router.GET("/:sid/json",ReturnJson)
	router.GET("/:sid/proto",ReturnProto)
	router.GET("/purejson", func(c *gin.Context) {
		c.Header("Content-Type","text/html;charset=utf-8")
		c.String(http.StatusOK,`<h1>Hello world!</h1>`)
	})
	router.POST("/signup",SignUpHandler)
	go func() {
		router.Run()
	}()
	quit:=make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGINT)
	<-quit
	log.Println("[ENDING] The Service Over!")
}
