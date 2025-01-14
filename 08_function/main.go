package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers1 := []int{3, 2, 3, 4, 5}
	dNumbers := transformNumbers(&numbers, tripleNumbers)
	fmt.Println(dNumbers)

	numbersFn2 := getTransformFn(&numbers1)

	dNumbers = transformNumbers(&numbers, numbersFn2)
	fmt.Println(dNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformFn(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return doubleNumbers
	} else {
		return tripleNumbers
	}
}

func doubleNumbers(number int) int {
	return number * 2
}
func tripleNumbers(number int) int {
	return number * 3
}
