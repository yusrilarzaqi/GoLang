package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Local())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Unix())

	utc := time.Date(2003, time.July, 23, 23, 59, 59, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, err := time.Parse(time.RFC3339, "2003-07-23T23:59:59Z")
	if err == nil {
		fmt.Println(parse)

	} else {
		fmt.Println(err.Error())
	}
}
