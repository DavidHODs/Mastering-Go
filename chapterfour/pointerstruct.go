package chapterfour

import "fmt"

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

func createStruct(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}

	return &myStructure{
		Name:    n,
		Surname: s,
		Height:  h,
	}
}

func reStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}

	return myStructure{
		Name:    n,
		Surname: s,
		Height:  h,
	}
}

func Struct() {
	s1 := createStruct("David", "HOD", 67)
	s2 := reStructure("David", "HOD", 182)

	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)
}