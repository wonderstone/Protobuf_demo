package god

import (
	"fmt"

	"github.com/wonderstone/protobuf_demo/proto/good"
)

// show01
func show01() {
	fmt.Println("Hello World")
	p := good.Person{
		Name:    "wonderstone",
		Age:     18,
		Address: "Shanghai",
	}
	g := God{
		
		Name:          "wonderstone",
		Age:           18,
		Persons:       []*good.Person{{Name: "wonderstone", Age: 18, Address: "Shanghai"},&p},
	}
	fmt.Println(g.Name)
}
