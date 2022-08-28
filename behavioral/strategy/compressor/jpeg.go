package compressor

import "fmt"

type jpec struct{}

func (j *jpec) Compress(data []byte) []byte {
	fmt.Println("Compressing JPEG")
	return data
}

func NewJpec() *jpec {
	return &jpec{}
}
