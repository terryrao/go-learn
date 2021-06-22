package go_memory_model

var bufferrdChannel1 = make(chan int, 0)
var a = ""

// A send on a channel happens before the corresponding receive from that channel completes.
func BufferredChannel1() {
	println("in gorountine")
	a = "hello world"
	bufferrdChannel1 <- 1
}

func helloInGogroutine() {
	go BufferredChannel1()
	<-bufferrdChannel1
	println(a)
	println("completed!")
}
