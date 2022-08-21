package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
  yusril := NewMap("Yusril Arzaqi")

  if yusril == nil {
    fmt.Println("Data Kosong")
  } else {
    fmt.Println(yusril)
  }
}

