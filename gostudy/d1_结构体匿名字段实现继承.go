package main
import "fmt"
type person struct {
	provice,city string
}
type student struct {
	name string
	age int
	person
}
func main(){
	//实例化方法1
	p1 :=person{"hunan","changsha"} 
	s1 :=student{"tom",15,p1}
	printinfo(s1)
	//实例化方法2
	s2 :=student{"zhangsan",28,person{"hubei","wuhan"}}
	printinfo(s2)

	//实例化方法3
	s3 :=student{}
	s3.name="jimmy"
	s3.age=30
	s3.provice="beijing"
    s3.city="beijing"

    printinfo(s3)

}

func printinfo(s student){
	fmt.Println(s)
	fmt.Printf("%+v \n",s)
	fmt.Printf("student name:%s,age:%d,provence:%s,city:%s\n", s.name,s.age,s.provice, s.city)
	fmt.Println("---------------------------------------")

}