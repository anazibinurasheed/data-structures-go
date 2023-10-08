package main

import "fmt"

// a recursive function will call itself directly or indirectly.

// it is used in divide and conquer algorithm.
// tree graph traversal
func main() {
	fmt.Println(Fibonacci(10))
	fmt.Println(IterativeFibonacci(10))
	fmt.Println(RecursiveFactorial(4))
	fmt.Println(IterativeFactorial(4))

}

// The Fibonacci numbers are the numbers in the following integer sequence: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ……..
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func IterativeFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	prev, current := 0, 1
	
	for i := 2; i <= n; i++ {
		next := (prev + current)
		prev = current
		current = next
	}

	return current

}

// n! =n*(n-1)!
// Factorial of a non-negative integer is the multiplication of all positive integers smaller than or equal to n. For example factorial of 6 is 6*5*4*3*2*1 which is 720.
func RecursiveFactorial(n int) int {
	//base case
	if n == 0 {
		return 1
	}

	fmt.Println("n", n, "n-1", n-1)
	return n * RecursiveFactorial(n-1)
}

func IterativeFactorial(n int) int {
	var res = 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}
