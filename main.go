package main

import "fmt"

const (
	BIAS = 132
	CLIP = 32635
)

func disp_c(value int){
	fmt.Printf("\t")
	for i:=0x80; i>0; i >>=1{
		if value & i != 0{}
	}
}

func g711encode(audio_sample int) int{
	var mask int

	if audio_sample > 0 && audio_sample > CLIP {audio_sample = CLIP}
	if audio_sample < 0 && audio_sample < -CLIP{audio_sample = -CLIP}
	if audio_sample < 0{
		mask = 0xff
		audio_sample = -audio_sample
	}else{
		mask = 0xff
	}
	audio_sample += BIAS
	i := 7
	for r:=0x4000; r>=0x80; r>>=1{
		b := r & audio_sample
		fmt.Println(b)
		if b != 0 {break}
		i--
	}
	mulaw_val := (i << 4) + ((audio_sample >> (i + 3)) & 0x0f)
	disp_c(mulaw_val)
	fmt.Println(mask)

	return mulaw_val
}



func main(){
	a_smp := []int{
		0, 140, 444, 1084, -2460, -5372, -11420, -24060,
	}
	fmt.Println(a_smp)

	t := 01

	fmt.Println(t)
	t <<= 2
	fmt.Println(t)


	g711encode(12)
}