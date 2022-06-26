package books

type BookType string

const (
	HardCover BookType = "HARD_COVER"
	SoftCover BookType = "SOFT_COVER"
	PaperBack BookType = "PAPER_BACK"
	EBook     BookType = "E_BOOK"
)

type Book struct {
	Name   string
	Author string
	Pages  int
	Type   BookType
}

type Library struct {
	Collection []Book
}

func (l *Library) IterateBooks(f func(Book) error) {
	for _, book := range l.Collection {
		if err := f(book); err != nil {
			break
		}
	}
}

func (l *Library) CreateIterator() iterator {
	return &BookIterator{
		books: l.Collection,
	}
}
