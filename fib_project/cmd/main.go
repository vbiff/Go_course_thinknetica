package main

import (
	fibonaci "fib_project/pkg"
	"flag"
	"fmt"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 0, "Введите число не больше 20")
	flag.Parse()

	if n > 20 {
		fmt.Println("Число больше 20")
		return
	}

	fmt.Println(fibonaci.Num(n))
}
