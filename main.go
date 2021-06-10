package main

import (
	"fmt"
)

const (
	BIAS = 132
	CLIP = 32635
)




func dispC(value int){

	fmt.Printf("%b", value)
}

func g711encode(audioSample int) int{
	var mask int

	if audioSample > 0 && audioSample > CLIP {
		audioSample = CLIP
	}else if audioSample < 0 && audioSample < -CLIP{
		audioSample = -CLIP
	}

	if audioSample < 0{
		mask = 0xff
		audioSample = -audioSample
	}else{
		mask = 0xff
	}
	audioSample += BIAS
	i := 7
	for r:=0x4000; r>=0x80; r>>=1{
		b := r & audioSample
		if b != 0 {break}
		i--
	}

	mulawVal := (i << 4) + ((audioSample >> (i + 3)) & 0x0f)
	dispC(mulawVal)
	return mulawVal ^ mask
}


func main(){
	aSmp := []int{
		0, 140, 444, 1084, -2460, -5372, -11420, -24060,
	}

	for a := range aSmp{
		muVal := g711encode(aSmp[a])
		fmt.Printf("\t%b", muVal)
		fmt.Printf("\n")
	}


}