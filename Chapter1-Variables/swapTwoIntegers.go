package swapTwoIntegers

import (
	"fmt"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	a, b = b, a

	fmt.Println(a, b)
}
