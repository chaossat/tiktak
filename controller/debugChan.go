package controller

var DebugChan = make(chan int, 10)

func DebugInit() {
	for i := 0; ; i++ {
		DebugChan <- i
	}
}
