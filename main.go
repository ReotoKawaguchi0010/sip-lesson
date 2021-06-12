package main

import (
	"fmt"
	"sio-lesson/codec"
)


func main(){
	aSmp := []int{
		0, 140, 444, 1084, -2460, -5372, -11420, -24060,
	}

	var eSmp []int

	for a := range aSmp{
		muVal := codec.G711encode(aSmp[a])
		eSmp = append(eSmp, muVal)
	}

	for e := range eSmp{
		d := codec.G711decode(eSmp[e])
		fmt.Println(d)
	}
}