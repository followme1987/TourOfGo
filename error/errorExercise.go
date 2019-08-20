package error

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot sqrt negative number", float64(e))
}

func Sqrt(number float64) (float64,error){
	if number < 0 {
		return 0, ErrNegativeSqrt(number)
	}else {
		return math.Sqrt(number),nil
	}

}
