package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println(*num)
}

func main() {
	num := 1

	fmt.Println("Before Change", num)

	changeNum(&num)

	fmt.Println("pointer num value after change", num)

}
