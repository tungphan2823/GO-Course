package main

import (
	"fmt"
)

func main() {
	age := 32

	agePointer := &age
	fmt.Printf("My age is %d years old.\n", *agePointer)
	getAdultYear(agePointer)
	fmt.Print(age)
}
func getAdultYear(age *int) {
	*age = *age - 18
}
