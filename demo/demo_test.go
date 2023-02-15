package demo

import (
	"fmt"
	// "github.com/golang/protobuf/proto"
	"testing"
)

// test demo.pb

func TestBarS(t *testing.T) {
	fmt.Println("Hello World")
	barS := BarS{
		BarTime: "2020-01-01 00:00:00",
		InstID:  "000001.SZ",
		IndiDataSlice: []float64{
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
		},
	}
	fmt.Println(barS.InstID)
	// write a demo protobuf message
	mapBarS := map[string]*BarS{
		"000001.SZ": &barS,
	}
	fmt.Println(mapBarS["000001.SZ"].InstID)
}
