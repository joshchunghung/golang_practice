package main

import "fmt"

func main() {
	//  basic
	// var x int = 0
	// for x < 3 {
	// 	fmt.Println(x)
	// 	x++
	// }

	//  basic2
	// var ii int
	// for ii = 0; ii <= 10; ii += 2 {
	// 	fmt.Println(ii)
	// }

	// 	整數求和
	var i, n, sum int
	fmt.Println("求和 1+2+....+n")
	fmt.Print("input n:")
	fmt.Scanln(&n)
	sum = 0
	for i = 0; i <= n; i++ {
		sum += i
	}
	fmt.Println("result is ", sum)
}
