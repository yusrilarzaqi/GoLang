package main

import (
	"fmt"
	"sort"
)

type User struct {
  Name string
  Age int
}

type UserSlide []User

func (userSlice UserSlide) Len() int {
  return len(userSlice)
}

func (userSlice UserSlide) Less(i, j int) bool {
  return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlide)  Swap(i, j int) { 
  userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
  users := UserSlide {
    {"Yusril", 18},
    {"Arzaqi", 18},
    {"Bimo", 17},
    {"Alamsyah", 17},
    {"Adam", 19},
  }

  fmt.Printf("users: %v\n", users)
  sort.Sort(users)
  fmt.Printf("users: %v\n", users)
}
