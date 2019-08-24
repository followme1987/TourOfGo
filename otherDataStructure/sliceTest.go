package otherDataStructure

import "fmt"

func SliceTest() {

	s := []struct{ x, y int }{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(s[0:1][0].x)

	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(s2[:2][1])
}

func SliceCapcity() {
	arr := [5]int{1, 2, 3, 4, 5}

	s := arr[:2]
	fmt.Printf("%d,%d,%v \n", len(s), cap(s), s)

	s = arr[1:4]
	fmt.Printf("%d,%d,%v \n", len(s), cap(s), s)

	s = arr[:]
	fmt.Printf("%d,%d,%v \n", len(s), cap(s), s)
	fmt.Printf("%d,%d,%v \n", len(arr), cap(arr), s)

	s = arr[2:3]
	fmt.Printf("%d,%d,%v \n", len(s), cap(s), s)

	var s2 []int
	if s2 == nil {
		fmt.Printf("%d,%d,%v \n", len(s2), cap(s2), s2)
	}
}

func SliceMake() {
	a := make([]int, 2)
	b := make([]int, 2, 4)

	fmt.Printf("%d,%d,%v \n", len(a), cap(a), a)
	fmt.Printf("%d,%d,%v \n", len(b), cap(b), b)
}

func SliceAppend() {
	var s []int
	s = append(s, 1)
	fmt.Println(s)

	s = append(s, 2, 3, 4, 5, 6)
	fmt.Println(s)
}

func SliceLoop() {

	s := []int{1, 2, 3, 4}
	for i, v := range s {
		fmt.Printf("%v,%v \n", i, v)
	}
}
