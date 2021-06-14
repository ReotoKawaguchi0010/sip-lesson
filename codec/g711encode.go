package codec

const (
	BIAS = 132
	CLIP = 32635
)



func G711encode(audioSample int) int{
	var mask int

	if audioSample > 0 && audioSample > CLIP {
		audioSample = CLIP
	}else if audioSample < 0 && audioSample < -CLIP{
		audioSample = -CLIP
	}

	if audioSample < 0{
		mask = 0x7f
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
	return mulawVal ^ mask
}

