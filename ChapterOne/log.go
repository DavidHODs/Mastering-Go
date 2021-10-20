package chapterone

// import (
// 	"log"
// 	"log/syslog"
// 	"os"
// 	"path/filepath"
// )

// // Logger demonstrates the logging capabilities that exists in Go
// func Logger() {
// 	programName := filepath.Base(os.Args[0])
// 	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)

// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.SetOutput(sysLog)
// 	}

// 	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")

// 	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program!")
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.SetOutput(sysLog)
// 	}

// 	log.Println("LOG_MAIL: Logging in Go!")
// 	log.Println("Will you see this?")
// }