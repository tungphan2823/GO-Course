package main

func main() {
	// list := []int{1, 3, 5, 6, 7}
	sum := sumUp(1, 3, 5, 6)
	println(sum)
}
func sumUp(list ...int) int {
	sum := 0
	for _, value := range list {
		sum += value
	}
	return sum
}
