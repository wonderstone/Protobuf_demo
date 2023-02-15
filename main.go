package main

import (
	"fmt"

	"github.com/wonderstone/protobuf_demo/demo"
	"github.com/wonderstone/protobuf_demo/proto/good"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("Hello World")
	// write a demo protobuf message
	barM := demo.BarM{
		BarTime: "2020-01-01 00:00:00",
		InstID:  "000001.SZ",
		IndiDataMap: map[string]float64{
			"open":   1.0,
			"high":   2.0,
			"low":    3.0,
			"close":  4.0,
			"volume": 5.0,
		},
	}
	fmt.Println(barM.BarTime)
	barC := demo.BarS{}
	barC.BarTime = "2020-01-01 00:00:00"
	barC.InstID = "000001.SZ"
	barC.IndiDataSlice = []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(barC.InstID)

	// encode, 转换成二进制数据
	data, err := proto.Marshal(&barC)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	// decode, 从二进制数据中解析出来
	barC2 := demo.BarS{}
	err = proto.Unmarshal(data, &barC2)
	if err != nil {
		panic(err)
	}
	fmt.Println(barC2.BarTime)

	g := good.Person{
		Name:    "wonderstone",
		Age:     18,
		Address: "Shanghai",
	}
	fmt.Println(g.Name)

}
