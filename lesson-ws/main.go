package main

import (
	"fmt"
	"sync"
)
type Student struct {
	Name string
	Age int
}
const N=10
func IsNil(i any) {
	if i == nil {
		fmt.Println("i is nil")
		return
	}
	fmt.Println("i isn't nil")
}

func main() {
	 m :=map[string]int{}
	m["a"]=5
	fmt.Println(m)
	var a []int
	a=append(a,9)
	fmt.Println(a)
	var sl []string
	if sl == nil {
	fmt.Println("sl is nil")
	}
	IsNil(sl)
	s:=&Student{}
	fmt.Println(s==nil)

	vector:=make(map[int]int)
	syncWG:=&sync.WaitGroup{}
	syncMU:=&sync.Mutex{}
	syncWG.Add(N)
	for i:=0;i<N;i++{
		go func() {
			defer syncWG.Done()
			syncMU.Lock()
			vector[i]=i
			syncMU.Unlock()
		}()
	}
	syncWG.Wait()
	fmt.Printf("%+v %d",vector,len(vector))

	array:=[]int{4,7,9}
	subArray:=array[1:]
	subArray=append(subArray,11)
	for _,v:= range subArray{
		v+=15
	}
	for i:=range subArray{
		subArray[i]+=15
	}
	fmt.Println(subArray)

	var arraya=make([]int,4,5)
	fmt.Printf("%d,%d,%v",len(arraya),cap(arraya),arraya)
	fmt.Println("abc"<"def")
	var maxSize int=5
vector1:=make(map[int]*int)
for index:=0;index<maxSize;index++{
	vector1[index]=&index
}
for _,value:=range vector1{
	fmt.Println(*value)
}

}
func f(){
	defer func() {
		fmt.Println("1;")
	}()
	defer func() {
		fmt.Println("2;")
	}()
	defer func() {
		fmt.Println("3;")
	}()
	panic("")
}