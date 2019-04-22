package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type JsonLog struct {
	TimeStamp  string
	LogMessage string
}

func main() {
	generate()
}

func generate() {
	for {
		generateSimpleLogLine()
		generateJsonLogLine()
		generateMultilineLogLine()
		generateErrorLogLine()

		time.Sleep(5 * time.Second)
	}
}

func generateSimpleLogLine() {
	fmt.Printf("[DockerLogGenerator] Current Time: %v\n", time.Now())
}

func generateJsonLogLine() {
	jsonMap := map[string]string{"timeStamp": fmt.Sprintf("%v", time.Now()), "logMessage": "[DockerLogGenerator] This is a JSON log entry"}
	json, _ := json.Marshal(jsonMap)
	fmt.Println(string(json))
}

func generateMultilineLogLine() {
	fmt.Printf("[DockerLogGenerator] Multiline: %v\n This is the second line\nThis is the third line\n", time.Now())
}

func generateErrorLogLine() {
	os.Stderr.WriteString("[DockerLogGenerator] This is an error, occured at " + fmt.Sprintf("%v", time.Now()))
}
