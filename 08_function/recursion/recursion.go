package recursion

import "fmt"

func main() {
	num := recursionFn(5)
	fmt.Println(num) // Output: 1200
}
func recursionFn(number int) int {
	if number == 0 {
		return 1
	}
	return number * recursionFn(number-1)
}
