package main

import "fmt"

func main() {
	var name = "Zidan"
	if name == "zidan" {
		fmt.Println(name)
	} else if name == "idan" {
		fmt.Println("Kurang Z mas")
	} else {
		fmt.Println("Ko salah yaa?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Kepanjangan gan")
	} else {
		fmt.Println("Jumlah huruf udah betul gan")
	}
}
