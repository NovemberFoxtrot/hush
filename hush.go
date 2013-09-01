package main

import (
	"fmt"
)

func activate(sum float64) float64 {
	if sum > 0 {
		return 1
	}

	return -1
}

func main() {
	inputs := [...]float64{12, 4}
	weights := [...]float64{0.5, -1}

	var sum float64

	for i := 0; i < len(inputs); i++ {
		sum += inputs[i] * weights[i]
	}

	output := activate(sum)

	fmt.Println(sum, output)
}
