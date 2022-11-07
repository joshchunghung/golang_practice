package main

import "fmt"

//  fmt.Println() 換行
//  fmt.Print() 不換行
func main() {
	//  print
	fmt.Print("print", 2, "sA", true)
	fmt.Println("println", 3, "s", false)
	// input
	// &變數名稱: 取得變數的指標(pointer)
	var x int
	fmt.Print("請輸入一個數字") // 不會換行	形成輸入的提示
	fmt.Scanln(&x)
	fmt.Println("x=", x)

	// 整合練習：輸入兩個整數，讓兩個整數相乘
	var x1 int
	var y1 int
	fmt.Print("input x=")
	fmt.Scanln(&x1)
	fmt.Print("input y=")
	fmt.Scanln(&y1)
	fmt.Println(x1 * y1)

	// 第二種寫法
	var x2 int
	var y2 int
	fmt.Print("input x y: ")
	fmt.Scanln(&x2, &y2)
	var result int = x2 * y2
	fmt.Println(result)
}
