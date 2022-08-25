package main

import (
	"flag"
	"fmt"
)

func main() {
	///////////////////(nama flag), (default), (description)
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
  var number = flag.Int("number", 0, "Put your number")

	flag.Parse()

  fmt.Println("host : ",*host)
  fmt.Println("Username : ",*username)
  fmt.Println("Password : ",*password)
  fmt.Println("number : ", *number)
}
