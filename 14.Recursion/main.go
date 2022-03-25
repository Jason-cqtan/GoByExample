package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fib2(n int) int {
	if n < 2 {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}

func main() {
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(28))
	fmt.Println(fib2(28))
}
