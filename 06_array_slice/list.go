package main

import (
	"fmt"
)

func main() {
	price := []float64{10.00, 8.99}
	fmt.Println(price)
	price = append(price, 5.99)
	fmt.Println(price)
}
