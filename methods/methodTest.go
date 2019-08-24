package methods

type Vex struct {
	X, Y int
}

type IntObj int

func (v Vex) MethodVex() int {
	return v.X + v.Y
}

func (i IntObj) MethodInt() int {
	return int(i)
}

func (v *Vex) MethodWithPointer() {
	v.X = 5
	v.Y = 6
}
