package main

import (
	"fmt"
)

type Vcard struct{
	name string
	address  *string 
	birth string
	imag  string
}

func prt_vc(cr *Vcard){

	fmt.Println("begin to print prt_vc")
	fmt.Println("--------------------------------")
	fmt.Printf("cr.name is :%s \n",cr.name)
	fmt.Printf("cr.address is :%s \n",*cr.address)
	fmt.Printf("cr.birth is :%s \n",cr.birth)
	fmt.Printf("cr.imag is :%s\n",cr.imag)
	fmt.Println("--------------------------------")
}


func main(){
	ad :="datang"
	var p *string 
	p=&ad
	vc :=&Vcard{"jonsen",p,"2018-5-1","image"}
	fmt.Printf("vc.name is :%v\n",vc)
	prt_vc(vc)
}

