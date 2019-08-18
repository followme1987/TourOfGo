package otherDataStructure

import "fmt"

func FunctionValue() {

	inner := func(x, y int) int {
		return x + y
	}

	fmt.Println(inner(1, 2))

	fmt.Println(Function2(inner))

	fmt.Println(Function2(func(x, y int) int {
		return x - y
	}))

	functionClosure := FunctionClosure()
	fmt.Println(functionClosure(4, 5))
}

func Function2(f func(int, int) int) int {
	return f(4, 5)
}

func FunctionClosure() func(int, int) int {
	return func(x, y int) int {
		return x * y
	}
}
