package main

import "fmt"

func changeValue(num *int) {
	*num = 30
}

func main() {

	num := 10

	fmt.Printf("Value of Num is %d \n", num)
	fmt.Println("changing value by passing it with refference")
	changeValue(&num)
	fmt.Printf("Value of Num is %d \n", num)

}
