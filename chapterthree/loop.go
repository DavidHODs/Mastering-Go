package chapterthree

import "fmt"

const (
	Zero int = iota
	One
	Two
	Three
	Four
)

const (
	p2_0 int = 1 << iota
	_
	p2_2
	_ 
	p2_4
	_ 
	p2_6
)

func Loop() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Println(i, " ")
	}

	fmt.Println()
	i := 10

	for {
		if i < 0 {
		break
		}
		fmt.Println(i, " ")
		i--
	}
	fmt.Println()

	i = 0
	anExpression := true

	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}

		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	anArray := [5]int{0, -1, 1, -2, 2}
	for i, value := range anArray {
		fmt.Println("index:", i, "value:", value)
	}

	// slice. If you add an element count e.g [20], it beomes an array
	// Slices are much more preferrable than arrays since they are more easier to manipulate  
	integer := make([]int, 20)

	for i := 0; i < len(integer); i++ {
		fmt.Println(integer[i])
	}

	fmt.Println()
	fmt.Println(Zero, One, Two, Three, Four)
	fmt.Println()
	fmt.Println(p2_0, p2_2, p2_4, p2_6)
}

