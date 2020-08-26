package maps

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func Maps() {
	// A map is a mapping of keys to values, like a dict in Python
	var myMap = make(map[string]Vertex)

	myMap["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(myMap)
	/*
		this will create a map like: [
			Bell Labs: {
				40.68433,
				-74.39967,1
			}
		]
	*/
}
