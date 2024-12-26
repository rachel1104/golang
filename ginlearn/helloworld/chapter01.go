package main

import "fmt"

func main() {
	key := ""
	loop := true
	fmt.Println("------------------------------------")
	fmt.Println("----------1.收支明细-----------------")
	fmt.Println("----------2.登记收入-----------------")
	fmt.Println("----------3.登记支出-----------------")
	fmt.Println("----------4.退出软件-----------------")
	fmt.Println("----------请选择1-4：-----------------")
	for {

		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("----------1.收支明细-----------------")
		case "2":
			fmt.Println("----------2.登记收入-----------------")
		case "3":
			fmt.Println("----------3.登记支出-----------------")
		case "4":
			loop = false
		default:
			fmt.Println("----------请输入正确选项-----------------")

		}
		if !loop {
			break
		}

	}
	fmt.Println("----------已退出-----------------")

}
