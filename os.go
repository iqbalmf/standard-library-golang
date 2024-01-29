package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
		//go run os.go "iqbal m fauzan"
		fmt.Println(arg)
	}
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	}
}
