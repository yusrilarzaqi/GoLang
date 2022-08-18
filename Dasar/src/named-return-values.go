package main

import "fmt"

func getComplateName() (firstName string, lastName string){
  firstName = "Yusril"
  lastName = "Arzaqi"

  return firstName, lastName
}

func main()  {
  firstName, lastName := getComplateName()
  fmt.Println(firstName, lastName)
}
