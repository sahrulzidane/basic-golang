package main

import "fmt"

func main() {
	var month = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// var slice1 = month[11:]
	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// month[5] = "diubah"
	// fmt.Println(slice1)

	// slice1[0] = "cek"
	// fmt.Println(slice1)

	// var slice2 = month[2:4]
	// fmt.Println(slice2)

	// var slice3 = append(slice2, "baru")
	// fmt.Println(slice3)
	// slice3[1] = "desember enakkk"
	// fmt.Println(slice3)
	// fmt.Println(month)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "zidan"
	newSlice[1] = "sahrul"
	fmt.Println(newSlice)

	copyslice := make([]string, len(newSlice), cap(newSlice))
	copy(copyslice, newSlice)
	fmt.Println(copyslice)
}
