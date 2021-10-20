package chapterone

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Error shows the error package in Go
func Error() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats")
		return
	}

	arguements := os.Args
	var err error = errors.New("an error just occured")

	k := 1
	var n float64

	for err != nil {
		if k >= len(arguements) {
			fmt.Println("None of the argument is a float!!")
			return
		}

		n, err = strconv.ParseFloat(arguements[k], 64)
		k++
	}

	min, max := n, n

	for i := 2; i < len(arguements); i++ {
		n, err := strconv.ParseFloat(arguements[i], 64)

		if err == nil {
			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}