package channelFibonacci

func FibWithChannel(ch chan int, size int) {

	x, y := 0, 1
	for i := 0; i < size; i++ {
		ch <- x
		x, y = y, x+y
	}

	close(ch)
}
