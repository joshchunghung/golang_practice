package main // 要可以執行一定是main
import "fmt"

func main() {
	// integer
	fmt.Println(3)
	// float64
	fmt.Println(3.1414)
	// string
	fmt.Println("SSS")
	// boolen
	fmt.Println(true)
	// rune : a : 97
	fmt.Println('a')
	fmt.Println("------------")
	// 變數
	var x int
	x = 4
	fmt.Println(x)
	x = 10
	fmt.Println(x)
	var y float64 = 2.33
	fmt.Println(y)
	var s string = "good"
	fmt.Println(s)
	var status bool = false
	fmt.Println(status)
	var d rune = 'b'
	fmt.Println(d)

}
