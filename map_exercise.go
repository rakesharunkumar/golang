package main

import "fmt"

func main() {
	x := map[string][]string{
		`personlist`: []string{`fp1,lp1`, `fp2,lp2`},
		`animallist`: []string{`al1,al2`, `al3,al4`},
		`birdlist`:   []string{`bl1,bl2`, `bl3,bl4`},
	}
	fmt.Println(x["personlist"])
	fmt.Println(x["animallist"])
	fmt.Println(x["birdlist"])
}
