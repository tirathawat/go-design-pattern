package main

import (
	"design-pattern/behavioral/strategy/compressor"
	"design-pattern/behavioral/strategy/filter"
	"design-pattern/behavioral/strategy/storage"
)

func main() {
	s := storage.New(compressor.NewJpec(), filter.NewBlackWhite())
	s.Store([]byte("Hello World"))

	s = storage.New(compressor.NewPng(), filter.NewBlackWhite())
	s.Store([]byte("Hello World"))

	s = storage.New(compressor.NewJpec(), filter.NewHighContrast())
	s.Store([]byte("Hello World"))

	s = storage.New(compressor.NewPng(), filter.NewHighContrast())
	s.Store([]byte("Hello World"))
}
