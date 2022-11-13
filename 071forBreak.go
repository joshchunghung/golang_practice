package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 10; i++ {
		if i == 5 {
			break // 直接結束迴圈
		}
		fmt.Println(i)
	}
}
