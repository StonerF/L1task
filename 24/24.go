package main

import (
	"fmt"
)

func main() {
	p1 := NewPoint(10.1, 15.2)
	p2 := NewPoint(-15.2, 3.1)

	line := Line{}
	fmt.Printf("distance between (10.1, 15.2) and (-15.2, 3.1) is %f\n", line.Distance(p1, p2))
}
