package chaptertwo

import (
	"fmt"
	"runtime"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)

	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("---------")
}

// PrintStats allowss the study of Go garbage collection operation
func Operation() {
	var mem runtime.MemStats

	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)

		if s == nil {
			fmt.Println("Operation Failed")
		}
	}

	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)

		if s == nil {
			fmt.Println("Operation Failed")
		}
	}

	printStats(mem)
}