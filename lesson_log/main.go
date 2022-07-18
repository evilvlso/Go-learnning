package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

func NormalUse() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	log.Println("this is a log after setflags")
}
func ToMultiOut() {
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("./log.log", os.O_WRONLY|os.O_CREATE, 0755)
	defer writer3.Close()
	if err != nil {
		log.Fatalf("crete file failed!")
	}
	logger := log.New(io.MultiWriter(writer1, writer2, writer3), "", log.Lshortfile|log.LstdFlags)
	logger.Printf("Customer logger!")
}

func main() {

	//log.Printf("this is a log")
	//NormalUse()
	ToMultiOut()
}
