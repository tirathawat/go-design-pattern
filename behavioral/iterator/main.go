package main

import (
	"design-pattern/behavioral/iterator/books"
	"fmt"
)

var lib *books.Library = &books.Library{
	Collection: []books.Book{
		{
			Name:   "War and Peace",
			Author: "Leo Tolstoy",
			Pages:  840,
			Type:   books.HardCover,
		},
		{
			Name:   "Crime and punishment",
			Author: "Leo Tolstoy",
			Pages:  1220,
			Type:   books.SoftCover,
		},
		{
			Name:   "Gan Banking",
			Author: "Gan Mongklakorn",
			Pages:  3000,
			Type:   books.PaperBack,
		},
		{
			Name:   "Sukitech",
			Author: "Sukitech.co.th",
			Pages:  10000,
			Type:   books.EBook,
		},
		{
			Name:   "M Maz",
			Author: "Sukitech.co.th",
			Pages:  10000,
			Type:   books.EBook,
		},
	},
}

func main() {
	lib.IterateBooks(func(book books.Book) error {
		fmt.Println("Book title:", book.Name)
		return nil
	})

	iter := lib.CreateIterator()
	for iter.HasNext() {
		fmt.Println("Book title:", iter.Next())
	}
}
