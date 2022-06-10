package Helper

import "fmt"

type Cat struct {
	name string
}

func (cat *Cat) Say() {
	fmt.Println("Meow")
}

func (cat *Cat) GetName() string {
	return cat.name
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

type Dog struct {
	name string
}

func (dog *Dog) Say() {
	fmt.Println("Awf")
}

func (dog *Dog) GetName() string {
	return dog.name
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}
