package protoDemo

import (
	"fmt"
	"github.com/wonderstone/protobuf_demo/proto/good"
)

func show() {
	fmt.Println("Hello World")
	p := good.Person{
		Name:    "wonderstone",
		Age:     18,
		Address: "Shanghai",
	}
	fmt.Println(p.Name)

}
