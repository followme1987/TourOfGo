package otherDataStructure

import "fmt"

type Vex struct {
	x, y int
}

var (
	v1 = Vex{x: 1}
	v2 = Vex{3, 5}
	v3 = Vex{}
	pp = &Vex{5, 6}
)

func PointerWithStruct() {
	fmt.Println(v1, v2, v3, pp)
}
func Pointers() {

	i := 10
	p := &i

	fmt.Println(p)
	fmt.Println(*p)

	*p = 12
	fmt.Println(i)
	fmt.Println(&i)

	vex := Vex{2, 3}
	v := &vex
	v.x = 12

	fmt.Println(vex.x)
}
