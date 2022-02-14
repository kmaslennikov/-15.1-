package main

import (
	"fmt"
)

func main() {
	var mass [10]int
	var countEven, countUneven int
	for i := 0; i < len(mass); i++ {
		fmt.Printf("Введите %v число: ", i+1)
		fmt.Scan(&mass[i])
	}
	for _, r := range mass {
		if r%2 == 0 {
			countEven++
		} else {
			countUneven++
		}
	}
	fmt.Printf("Четных: %v, нечетных: %v\n", countEven, countUneven)

}
