package consumer

import (
	"flag"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
)

func GetRectangleAndSquareArea(address string) ([2]float32, error) {
	var a [2]float32
	a[0] = 12
	a[1] = 9
	return a, nil
}
