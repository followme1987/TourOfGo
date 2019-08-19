package typeSwitch

import "fmt"

func Do(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("it is a integer")
	case string:
		fmt.Println("it is a string")
	default:
		fmt.Println("I do not know")
	}
}

func CheckType() {
	var t interface{}
	t = 10
	v := t.(bool)

	fmt.Println(v)


}