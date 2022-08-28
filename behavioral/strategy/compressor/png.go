package compressor

import "fmt"

type png struct{}

func (j *png) Compress(data []byte) []byte {
	fmt.Println("Compressing PNG")
	return data
}

func NewPng() *png {
	return &png{}
}
