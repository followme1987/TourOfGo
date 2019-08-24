package interfaceTest

import "fmt"

type Human interface {
	Eat() string
}

func (stu *Student) Eat() string {
	return "student can eat"
}

type Student struct {
	Name string
}

func EmptyInterfaceTest() {
	var anything interface{}

	anything = "string"
	fmt.Println(anything)
	anything = 2
	fmt.Println(anything)
}
