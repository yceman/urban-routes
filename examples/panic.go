package panic

import "fmt"

func divide(a, b int) int {
	result := a / b
	return result
}

func panic() {
	fmt.Println(divide(1, 0))
	fmt.Println(divide(1, 1))
}