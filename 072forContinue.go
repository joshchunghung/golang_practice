package main

import "fmt"

func main() {
	var i int
	for i = 0; i <= 10; i++ {
		if i == 5 {
			continue //跳過本次回圈
		}
		fmt.Println(i)
	}
}
