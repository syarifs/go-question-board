package logger

import (
	"fmt"
	"log"
	"os"
)

type Logging struct {}
var Logger = log.New(Logging{}, "", log.Lmicroseconds|log.Lshortfile)

func (e Logging) Write(p []byte) (n int, err error) {
	fmt.Println("Error: " + string(p))
	file, _ := os.OpenFile("question-board.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	file.WriteString(string(p))

	// Close the file when the surrounding function exists
	defer file.Close()

	return n, err
}

func WriteLog(logs interface{}) {
	Logger.Println(logs)
}

