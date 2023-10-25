package main

import "fmt"

func main() {
	var names [8]string
	names[0] = "zidane"
	names[1] = "sahrul"
	names[2] = "romadon"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(len(names))

	var values = [3]int{
		12,
		31,
		14,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

}
