package main

import "fmt"

func main() {
	//  basic if
	// if false {
	// 	fmt.Println("go!!")
	// } else {
	// 	fmt.Println("NO!!!!")
	// }

	// 判斷輸入值是否適當
	var money int
	fmt.Print("請輸入金額：")
	fmt.Scanln(&money)

	if money < 0 {
		fmt.Println("input wrong!")
	} else if money <= 20000 {
		fmt.Println("OK!!")
	} else {
		fmt.Println("Too Much!")
	}

}
