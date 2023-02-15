package bad

import (
	"fmt"
)

func show01() {
	fmt.Println("Hello World")

	d := Dog{
		Name:    "dog",
		Age:     1,
		Address: "Shanghai",
	}
	p := PersonBad{
		Name:    "wonderstone",
		Age:     18,
		Address: "Shanghai",
		Dogs:    []*Dog{&d},
	}
	fmt.Println(p.Name)
}
