package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int8	`json:"age"`
	School School
	Degree string 	`json:"degree"`
}
type School struct {
	SchoolName string 	`json:"school_name"`
	Level int8	`json:"level"`
}
func main() {
	s:=School{"西安交通大学",1}
	p:=Person{"李逵",28,s,"Master"}
	fmt.Printf("%T---%v\n",p,p)
	j,err:=json.Marshal(p)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("serialize result:%T---%v\n",j,string(j))
	// 反序列化
	var p1 Person
	err=json.Unmarshal(j,&p1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("un serialize result:%T---%v\n",p1,p1)
}