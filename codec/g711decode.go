package codec

func G711decode(muVal int) int {
	var audioSample int
	var t int

	ulawSample := ^muVal
	t = ((ulawSample & 0x0f) << 3) + BIAS
	t <<= (ulawSample & 0x70) >> 4
	if ulawSample & 0x80 > 0 {
		audioSample = BIAS - t
	}else{
		audioSample = t - BIAS
	}

	return audioSample
}
