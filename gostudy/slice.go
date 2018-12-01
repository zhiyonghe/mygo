package main 

import "fmt"



func main (){
	slic :=[] int{1,30,50,24,35,13,34}
	temp :=[]int {}
	fmt.Println("slice len is :",len(slic))
	sl := slic[2:5]
	fmt.Println("s1 len is :",len(sl))
	for i,v := range sl {
		fmt.Printf("index i is:%d,value is :%d\n", i,v)
		// fmt.Println(val)
		temp=append(temp,v)
		temp=append(temp,v)
		}	
	fmt.Println(temp)
	fmt.Println(slic)
	
	fuirt :=[] string{"apple","banana","orange","grape","plum"}
	f1 :=fuirt[1:3:3] // 加上第三个引索值后，在继续append 相当于f1 有自己的底层数组
	fmt.Println("befor append f1:",f1)
	for _, k :=range fuirt{
		f1=append(f1,k)

	}
	fmt.Println(f1)
}