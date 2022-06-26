package books

type IterableCollection interface {
	CreateIterator() iterator
}

type iterator interface {
	HasNext() bool
	Next() *Book
}

type BookIterator struct {
	current int
	books   []Book
}

func (i *BookIterator) HasNext() bool {
	if i.current < len(i.books) {
		return true
	}

	return false
}

func (i *BookIterator) Next() *Book {
	if i.HasNext() {
		book := i.books[i.current]
		i.current++
		return &book
	}

	return nil
}
