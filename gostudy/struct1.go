package main
import (
	"fmt"
)

type struct1 struct{

	id int
	fl float32
	str string
}

func test_strc(st *struct1){
   fmt.Println("-------begin test strunct---------------")
   fmt.Printf("st.id is: %d \n",st.id)
   fmt.Printf("st.fl is: %f \n",st.fl)
   fmt.Printf("st.str is: %s \n",st.str)
   fmt.Println("-------end test strunct---------------")
}

func main(){
	msg :=new(struct1)
	msg.id=10
	msg.fl=10.5
	msg.str="zhangsan"
	fmt.Printf("msg.id is: %d \n",msg.id)
	fmt.Printf("msg.fl is: %f \n",msg.fl)
	fmt.Printf("msg.str is: %s \n",msg.str)
	fmt.Println(msg)
	var  k struct1    //定义一个strunct1变量
	k.id=20
	k.fl=1.9
	k.str="test"
	fmt.Printf("k.id is: %d \n",k.id)
	fmt.Printf("k.fl is: %f \n",k.fl)
	fmt.Printf("k.str is: %s \n",k.str)
	fmt.Println(k)

	var p *struct1 
	p=msg
	// p.id=20
	// p.fl=30.9
	// p.str="hehe"
	fmt.Printf("p.id is: %d \n",p.id)
	fmt.Printf("p.fl is: %f \n",p.fl)
	fmt.Printf("p.str is: %s \n",p.str)
	fmt.Println(p)
	fmt.Println(*p)

	msg2 :=new(struct1)    //new 出来的就是一个指针地址
	*msg2=struct1{40,40.5,"good boy"}
	fmt.Println("msg2 is ",*msg2)
	fmt.Printf("msg2.id is :%d\n",msg2.id)


	test_strc(p)
	test_strc(&k)

}