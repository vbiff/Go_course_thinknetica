package fibonaci

import (
	"fmt"
)

// Fibo - fibonaci counter
func Fibo(n int) int {
	if n > 20 {
		fmt.Println("Число не должно превышать 20, вы ввели", n)
		return -1
	}
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return Fibo(n-1) + Fibo(n-2)
}
