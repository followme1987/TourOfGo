package main

import (
	"fmt"
	"github.com/followme1987/TourOfGo/controlFlow"
	"github.com/followme1987/TourOfGo/error"
	"github.com/followme1987/TourOfGo/interfaceTest"
	"github.com/followme1987/TourOfGo/methods"
	"github.com/followme1987/TourOfGo/otherDataStructure"
	"github.com/followme1987/TourOfGo/reader"
	"github.com/followme1987/TourOfGo/stringers"
	"github.com/followme1987/TourOfGo/typeSwitch"
	"io"
	"math"
	"math/rand"
	"os"
	"strings"
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

	sum := controlFlow.ForLoop()

	fmt.Println(sum)

	sum2 := controlFlow.ForLikeWhile()
	fmt.Println(sum2)

	controlFlow.IfStatement()

	controlFlow.SwitchStatement()

	controlFlow.HowFarFromSat()

	controlFlow.SwitchWithout()

	controlFlow.DeferStatement()

	controlFlow.StackDefer()

	otherDataStructure.Pointers()

	otherDataStructure.PointerWithStruct()

	otherDataStructure.ArrayTest()

	otherDataStructure.SliceTest()

	otherDataStructure.SliceCapcity()

	otherDataStructure.SliceMake()

	otherDataStructure.SliceAppend()

	otherDataStructure.SliceLoop()

	otherDataStructure.WordCounter("foo boo zoo you foo foo you")

	otherDataStructure.FunctionValue()

	otherDataStructure.Fib()

	v := methods.Vex{1, 2}
	fmt.Println(v.MethodVex())

	i := methods.IntObj(2)
	fmt.Println(i.MethodInt())

	v2 := &methods.Vex{1,2} //&Methods.Vex{1,2} will work as well that is the tradeoff of pointer receiver
	v2.MethodWithPointer()
	fmt.Println(v2.X,v2.Y)

	fmt.Println(v2.MethodVex())

	var human interfaceTest.Human
	stu := &interfaceTest.Student{"tom"}
	human = stu
	fmt.Println(human.Eat())

	interfaceTest.EmptyInterfaceTest()

	CheckType()

	typeSwitch.Do("10")

	ip := stringers.IPAddr{127,0,0,1}
	fmt.Println(ip)

	if v,err:=error.Sqrt(2); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(v)
	}

	if v,err:=error.Sqrt(-2); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(v)
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := reader.Rot13Reader{s}
	io.Copy(os.Stdout, &r)
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

	fmt.Printf("the data type is %T,%T \n", num, num2)
}
