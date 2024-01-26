package algorithms

import "fmt"

func extractNumbersFromAnInteger(num int) {
	digit := 0
	for num > 0 {
		digit = num % 10
		fmt.Println(digit)
		num = num / 10
	}

}
