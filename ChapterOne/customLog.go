package chapterone

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "tmp/mGo.log"

// CustomLog writes log info to a custom file
func CustomLog() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)

	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello There AgainIIII!")
	iLog.Println("Another Log Entry")
}