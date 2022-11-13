package main

import "fmt"

// parameter : 參數名稱 資料型態
func show(msg string, i int) {
	fmt.Println(msg, i)
}

// 有回傳值
func add(x float64, y float64) float64 {
	return x + y
}

func main() {
	show("hellp", 123)
	var addResult float64
	addResult = add(2.0, 5.0)
	fmt.Println(addResult)
}
