package main

import (
	"fmt"
	"net/rpc"
	"os"
)

type Result struct{
	Num int
	DoubleNum int
}

func main() {
	quit := make(chan os.Signal,1)
	//result:=&Result{}
	var result Result
	client,err:=rpc.DialHTTP("tcp","127.0.0.1:8899")
	if err != nil {
		fmt.Println("Dial error")
		fmt.Println(err)
	}
	if err=client.Call("Square.s",12,&result);err != nil {
		fmt.Println("Call error")
		fmt.Println(err)
	}
	fmt.Printf("RPC result %v",result)
	<-quit
}
