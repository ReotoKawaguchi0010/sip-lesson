package main

import "fmt"


const (
	BIAS = 132
	CLIP = 32635
)

func g711encode(audio_sample int) string{
	var _range int
	var mask int
	var mulaw_val string

	if audio_sample > 0 && audio_sample > CLIP {audio_sample = CLIP}
	if audio_sample < 0 && audio_sample < -CLIP{audio_sample = -CLIP}
	if audio_sample < 0{
		mask = 0xff
		audio_sample = -audio_sample
	}else{
		mask = 0xff
	}
	fmt.Println(mask)
	audio_sample += BIAS
	for i:=7; _range >= 0x80; i--{}
	return mulaw_val
}



func main(){
	a_smp := []int{
		0, 140, 444, 1084, -2460, -5372, -11420, -24060,
	}
	fmt.Println(a_smp)
}