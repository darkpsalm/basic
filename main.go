
package main

func main() {
	i()
}

func i() chan int {
	return make(chan int)
}
