package filter

import "fmt"

type blackWhite struct{}

func (b *blackWhite) Filter(data []byte) []byte {
	fmt.Println("Filtering Black and White")
	return data
}

func NewBlackWhite() *blackWhite {
	return &blackWhite{}
}
