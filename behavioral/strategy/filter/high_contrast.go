package filter

import "fmt"

type highContrast struct{}

func (h *highContrast) Filter(data []byte) []byte {
	fmt.Println("Filtering High Contrast")
	return data
}

func NewHighContrast() *highContrast {
	return &highContrast{}
}
