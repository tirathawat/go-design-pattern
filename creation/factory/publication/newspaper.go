package publication

import "fmt"

type newspaper struct {
	publication
}

func (n newspaper) String() string {
	return fmt.Sprintf("Newspaper\nName: %s\nPages: %d\nPublisher: %s\n",
		n.publication.name,
		n.publication.pages,
		n.publication.publisher,
	)
}

func createNewspaper(name, publisher string, pages int) *newspaper {
	return &newspaper{
		publication: publication{
			name:      name,
			publisher: publisher,
			pages:     pages,
		},
	}
}
