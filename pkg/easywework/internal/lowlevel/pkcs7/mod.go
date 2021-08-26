package pkcs7

func Pad(x []byte) []byte {
	numPadBytes := 32 - len(x)%32
	padByte := byte(numPadBytes)
	tmp := make([]byte, len(x)+numPadBytes)
	copy(tmp, x)
	for i := 0; i < numPadBytes; i++ {
		tmp[len(x)+i] = padByte
	}
	return tmp
}

func Unpad(x []byte) []byte {
	// last byte is number of suffix bytes to remove
	n := int(x[len(x)-1])
	return x[:len(x)-n]
}
