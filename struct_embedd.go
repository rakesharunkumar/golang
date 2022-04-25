package main

import (
	"fmt"
	//"image/color"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(truck1, sedan1)
	fmt.Println(truck1.fourWheel, sedan1.luxury)
}
