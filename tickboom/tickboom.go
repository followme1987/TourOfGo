package tickboom

import (
	"fmt"
	"time"
)

func TickBoom() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("  .")
			time.Sleep(50 * time.Millisecond)
		}

	}
}
