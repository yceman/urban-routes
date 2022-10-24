package main

import (
	"fmt"

	routes "github.com/yceman/urban-routes/application"
)

func main() {
	route := routes.Routes{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
	//fmt.Println("Hello 世界")
}
