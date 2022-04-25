package main

import (
	"fmt"
)

func main() {
	maps := map[string]string{"firstname": "Romeo", "Lastname": "Juliet"}
	fmt.Println(maps)
	for k, v := range maps {
		fmt.Printf("%s\t Is My %s\n", v, k)
		//fmt.Println(k, v)
	}
	// append map
	maps["Country"] = "India"
	fmt.Printf("Appended list:")
	fmt.Println(maps)
	fmt.Println(maps["Country"])
}
