package main

import (
	"flag"
	"fmt"
)

func main()  {
	//package flag berisikan fungsionalitas untuk memparsing command line argument
	host := flag.String("host", "localhost", "put your database host")
	username := flag.String("username", "root", "put your database username")
	password := flag.String("password", "root", "put your database password")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()
	
	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("host", *host)
	fmt.Println("port", *port)
	/* untuk run
	* go run flag.go -username=iqbal -password="rahasia banget" -host=123.123.23.3 -port=5505
	*/
}