package main

import (
	"fmt"
	"thirthfamous/possibility-hour/services"
)

func main() {
	// Example test cases.
	fmt.Println(services.PossibleTimes([]int{1, 2, 3, 4}))
	fmt.Println(services.PossibleTimes([]int{9, 1, 2, 0}))
	fmt.Println(services.PossibleTimes([]int{2, 2, 1, 9}))
}
