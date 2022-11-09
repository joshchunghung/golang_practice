package main

import (
	"fmt"
	"math"
)

func main() {
	// ##### 一般運算 + - * / %
	//  整數除法
	fmt.Println("整數除法")
	var x1 int = 10
	var y1 int
	y1 = x1 / 3
	fmt.Println("y1=10/3 , y1=", y1)
	fmt.Println("小數除法")
	var x2 float64 = 10.0
	var y2 float64
	y2 = x2 / 3.
	fmt.Println("y2=10/3 , y2=", y2)
	// 保留小数以fmt.Sprintf来实现，转为字符串后再以strconv.ParseFloat来转为浮点数。
	fmt.Println(fmt.Sprintf("%.2f", y2))

	//  次方 pow
	var y = math.Pow(2, 0.5) // 2**3 = 8
	fmt.Println(y)
	// 指定運算 =,+=.-=,*=,/=
	var x3 int = 2
	x3 += 2 // x3=x3+2
	fmt.Println(x3)

	// 單元運算 ++ --
	var x4 int = 4
	x4-- // x4=x4-1

	// 比較運算 >,<,>=,<=,==
	var status bool = 3 == 3
	fmt.Println(status)

	// 邏輯運算 ! && ||
	status = !status
	fmt.Println(status)
	status = (!false && true) || false
	fmt.Println(status)
}
