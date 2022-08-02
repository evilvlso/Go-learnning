package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	pd "github.com/mprotoc/pd"
	_ "io/ioutil"
)

func main() {
	sr:=pd.SumRequest{Num: 2}
	data, err :=proto.Marshal(&sr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v,%T",data,data)
}