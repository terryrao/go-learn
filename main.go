package main

var c = make(chan int, 10)
var a string

func f() {
	println("in gorountine")

	a = "hello, world"
	<-c
}

func main() {
	go f()
	c <- 0
	print(a)
}
