package controlFlow

import (
	"fmt"
	"runtime"
	"time"
)

func ForLoop() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func ForLikeWhile() int {
	sum := 1
	for sum < 10 {
		sum += 1
	}
	return sum
}

func IfStatement() {
	if condition := ForLikeWhile(); condition == 10 {
		fmt.Printf("if statement is in %v \n", condition)
	} else {
		fmt.Println("not in ")
	}
}

func SwitchStatement() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("mac os")
	case "linux":
		fmt.Println("linux os")
	default:
		fmt.Println("something else")
	}
}

func HowFarFromSat() {

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in 2 days")
	default:
		fmt.Println("too far")
	}
}

func SwitchWithout() {
	time := time.Now().Hour()
	switch {
	case time < 12:
		fmt.Println("good morning")
	case time < 17:
		fmt.Println("good evening")
	default:
		fmt.Println("hello")
	}
}

func DeferStatement() {
	fmt.Println("hello")
	defer fmt.Println("YI")
	fmt.Println("wu")
	fmt.Println("che")

	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
}

func StackDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
