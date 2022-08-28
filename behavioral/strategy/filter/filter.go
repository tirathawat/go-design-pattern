package filter

type Filter interface {
	Filter(data []byte) []byte
}
