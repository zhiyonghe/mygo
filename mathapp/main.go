package main
import (
	"time"
	s  "mygo/mymath"
	"fmt"
	std  "mygo/gostudy"
)
const LIM=41           // 定义常量
var fibs [LIM] uint64  //定义数组
func loop(){
    for i :=0;i<10;i++ {

     fmt.Printf("%d,",i)
		}

}
func fibonacci(n int)(res uint64){
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

func test_fib(){
	var result uint64=0
	start := time.Now()
	for i:=0;i<LIM;i++{
		result=fibonacci(i)
		fmt.Printf("fibonacci(%d) is:%d\n",i,result)

	}
	end := time.Now()
	delta :=end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)



}

func main() {
		fmt.Printf("Hello, world.  Sqrt(4) = %v\n", s.Sqrt(4))
		//go loop()
		//loop()
		// fmt.Printf("constr:%d\n" ,LIM)va
		// fmt.Printf("var:%d\n",fibs)
		time.Sleep(time.Second) 
		test_fib()
		//test  defer function
		std.Test_defer_fc()
		std.Trace_defer()
}