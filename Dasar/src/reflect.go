package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	// Name string
	Name string `required:"true" max:"10"`
}

type Sample1 struct {
  Name, desc string `required:"true"`
}

/*
@parameter data any
@return boolean
*/
func IsValid(data any) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			// return reflect.ValueOf(data).Field(i).Interface() != ""
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			} else {
				return true
			}
		}
	}

	return true
}

func main() {
	sample := Sample{"Yusril"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	// required := structField.Tag.Get("required")

	fmt.Printf("sample: %v\n", sample)
	fmt.Println("sampleType", sampleType)
	fmt.Println(structField.Type)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
  sample.Name = ""
	fmt.Println(IsValid(sample))


  contoh := Sample1{"", ""}
  fmt.Println(IsValid(contoh))
  
}

