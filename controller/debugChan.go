package controller

var DebugChan = make(chan int, 10)

//DebugInit:初始化debug
func DebugInit() {
	for i := 0; ; i++ {
		DebugChan <- i
	}
}
