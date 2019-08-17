package main

import (
	"fmt"
	"math"
	"math/rand"
)

var (
	python, c = true, false
	cplus     = "Yes"
)

func main() {
	const Key = "this is a const value \n"

	fmt.Printf(Key)

	fmt.Printf("hello，世界 \n")
	fmt.Println("my faviour time is :", rand.Intn(10))
	fmt.Println(add(1, 2))
	wordA, wordB := swap("hello", "world")
	fmt.Println(wordA, wordB)
	fmt.Println(split(16))

	var java string = "no"

	csharp := "yes"
	fmt.Println(python, c, cplus, java, csharp)

	Conversion()

	CheckType()
}

func add(x, y int) int {
	return x + y
}

// return Multi results
func swap(wordA, wordB string) (string, string) {
	return wordB, wordA
}

//named return value
func split(sum int) (x, y int) {
	x = sum - 4
	y = sum * 16
	return
}

//type conversion
func Conversion() {
	x, y := 1, 2
	f := math.Sqrt(float64(x*x + y*y))
	u := uint(f)

	fmt.Printf("%v, %v\n", f, u)
}

//check type
func CheckType() {
	num := 1
	num2 := 1.1

	fmt.Printf("the data type is %T,%T", num, num2)
}
