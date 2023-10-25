package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "zidan",
		"address": "bekasi",
	}
	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "zidane"
	book["wrong"] = "ups"
	fmt.Println(book["title"])
	book["wrong"] = "haha"
	fmt.Println(book["wrong"])
	fmt.Println(len(book))

	delete(book, "wrong")

	fmt.Println(len(book))

}
