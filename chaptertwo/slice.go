package chaptertwo

import "runtime"

type data struct {
	i, j int
}

// Slice is used for storing large number of data structures
func Slice() {
	var N = 40000000
	var structure []data

	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(
			structure, data{value, value})
	}

	runtime.GC()
	_ = structure[0]
}

// MapPointres stores values with pointers
func MapPointers() {
	var N = 40000000
	myMap := make(map[int]*int)

	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = &value
	}

	runtime.GC()
	_ = myMap[0]
}

// Map stores plain values without pointers
func Map() {
	var N = 40000000
	myMap := make(map[int]int)

	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = value
	}

	runtime.GC()
	_ = myMap[0]
}

// Split splits maps into map of maps
func Split() {
	var N = 40000000
	split := make([]map[int]int, 200)

	for i := range split {
		split[i] = make(map[int]int)
	}

	for i := 0; i < N; i++ {
		value := int(i)
		split[i%200][value] = value
	}

	runtime.GC()
	_ = split[0][0]
}