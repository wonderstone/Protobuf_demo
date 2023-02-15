package good

import (
	"fmt"
)

func show() {
	fmt.Println("Hello World")
	p := Person{
		Name:    "wonderstone",
		Age:     18,
		Address: "Shanghai",
	}
	fmt.Println(p.Name)
}
