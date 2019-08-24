package stringers

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {

	str := fmt.Sprintf("%v", ip[0])
	for _, v := range ip[1:] {
		str += fmt.Sprintf(".%v", v)
	}
	return fmt.Sprintf("%v", str)
}
