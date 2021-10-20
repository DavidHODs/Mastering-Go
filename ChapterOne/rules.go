package chapterone

// 1. You cannot import a package and then leave it unused
// 2. Putting a _ before an unused Go package will make Go compiler bypass this error
import (
	"fmt"
	_ "os"
)

// Rules simply demonstrates the rules guiding unused Go packages
func Rules() {
	fmt.Println("Hello There!")
}