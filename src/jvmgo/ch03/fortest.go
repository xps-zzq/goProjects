package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	y := 1
	for ; y < 1000; {
		y += y
	}
	fmt.Println(y)
	//while
	x:=0
	for x<10  {
		x+= 1
	}
	fmt.Println(x)

	for {
		fmt.Println("loop")
		break
	}
}
