package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
  // Person
	{
		yusril := Person{Name: "Yusril Arzaqi"}
		myName := yusril.GetName()
		SayHello(yusril)
		fmt.Println(myName)
	}

  // Animal
  {
    dog := Animal{Name: "Dog"}
    fmt.Println(dog.GetName())
    SayHello(dog)
  }
}
