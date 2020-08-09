package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	f,err :=os.Open("e://test.ini")
	if err !=nil {
		fmt.Println("os open file err",err)
		return

	defer f.Close()
	fmt.Println("begin to read file.")
	b,err :=ioutil.ReadAll(f)
	if err !=nil{
		fmt.Println("ioutil ReadAll error: ", err)
		return
	}
	fmt.Println(b)
	}

}