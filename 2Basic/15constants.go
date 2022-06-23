package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "Hello"
	fmt.Println("hello world")
	fmt.Println("happy ", Pi, "day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}