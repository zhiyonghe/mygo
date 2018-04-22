
package gostudy
import (
	"fmt"
)


func fct1(){
	fmt.Printf("test 1 print first\n")
	defer fct3()
	fmt.Printf("test 1 print second\n")
	defer fct2()
 }
 
 func fct2(){
	fmt.Printf("test 2 print first\n")
 }
 
 func fct3(){
	fmt.Printf("test 3 print first\n")
 }
func Test_defer_fc(){
   fmt.Printf("begin to test func defer")
   fct1() 
   fmt.Printf("end test func defer\n")
 
}

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
    trace("a")
    defer untrace("a")
    fmt.Println("in a")
}

func b() {
    trace("b")
    defer untrace("b")
    fmt.Println("in b")
    a()
}

func Trace_defer(){
	b()
}