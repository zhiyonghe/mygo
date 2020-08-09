package main

import (
	"fmt"
	"math"

)
type Point3 struct{ x,y,z float64}

func (p *Point3) Abs() float64{

	return math.Sqrt(p.x*p.x+p.y*p.y+p.z*p.z)
}


func main(){

	t :=new(Point3)      // 实例化
	t.x=3                //选择器
	t.y=4
	t.z=5
	fmt.Println(t.Abs())

	k :=&Point3{8,9,10}  //实例化，并初始化值
	fmt.Println(k.Abs())

}