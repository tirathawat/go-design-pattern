package storage

import (
	"design-pattern/behavioral/strategy/compressor"
	"design-pattern/behavioral/strategy/filter"
)

type storage struct {
	compressor compressor.Compressor
	filter     filter.Filter
}

func New(compressor compressor.Compressor, filter filter.Filter) *storage {
	return &storage{
		compressor: compressor,
		filter:     filter,
	}
}

func (s *storage) Store(data []byte) []byte {
	data = s.filter.Filter(data)
	data = s.compressor.Compress(data)
	return data
}
