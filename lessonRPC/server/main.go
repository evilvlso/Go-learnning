package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type  Result struct{
	Num int
	DoubleNum int
}
type Cal int
func (c *Cal) Square(n int,result *Result) error  {
	result.Num=n
	result.DoubleNum=n*n
	log.Printf("the param is %d",n)
	return nil
}

func main() {
rpc.RegisterName("Square.s",new(Cal))
rpc.HandleHTTP()
log.Printf("the rpc server serving on port 8899")
if err:=http.ListenAndServe(":8899",nil);err!=nil{
	log.Fatal("Error serving: ", err)
}
}
