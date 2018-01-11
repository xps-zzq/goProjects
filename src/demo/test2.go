package main

import "fmt"


type Vertexx struct {
	X, Y int
}

var (
	p = Vertexx{1, 2}  // 类型为 Vertex
	q = &Vertexx{1, 2} // 类型为 *Vertex
	r = Vertexx{X: 1}  // Y:0 被省略
	s = Vertexx{}      // X:0 和 Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}
