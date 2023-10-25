package main

import "fmt"

func main() {
	const device = "laptop"
	fmt.Println(device)

	const (
		alis string = "pensil"
		pake        = "kertas"
	)
	fmt.Println(alis, pake)
}
