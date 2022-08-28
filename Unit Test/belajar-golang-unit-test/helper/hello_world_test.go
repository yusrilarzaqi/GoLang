package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldYusril(t *testing.T) {
	result := HelloWorld("Yusril")

	if result != "Hello Yusril" {
		// unit test failed
		t.Fail()
	}

	fmt.Println("TestHelloWorldYusril done")
}

func TestHelloWorldArzaqi(t *testing.T) {
	result := HelloWorld("Arzaqi")

	if result != "Hello Arzaqi" {
		// unit test failed
		t.FailNow()
	}

	fmt.Println("TestHelloWorldArzaqi done")
}

func TestHelloWorldBimo(t *testing.T) {
	result := HelloWorld("Bimo")

	if result != "Hello Bimo" {
		t.Error("Harusnya Hello Bimo")
	}

	fmt.Println("Dieksekusi")
}

func TestHelloWorldAlamsyah(t *testing.T) {
	result := HelloWorld("Alamsyah")

	if result != "Hello Alamsyah" {
		t.Fatal("Harusnya Hello Alamsyah")
	}

	fmt.Println("Tidak Dieksekusi")
}
