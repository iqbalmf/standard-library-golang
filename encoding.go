package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var encoding = base64.StdEncoding.EncodeToString([]byte("Iqbal M Fauzan"))
	fmt.Println(encoding)
	decode, err := base64.StdEncoding.DecodeString(encoding)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(string(decode))
	}
}
