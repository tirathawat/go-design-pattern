package publication

import "fmt"

type magazine struct {
	publication
}

func (n magazine) String() string {
	return fmt.Sprintf("Magazine\nName: %s\nPages: %d\nPublisher: %s\n",
		n.publication.name,
		n.publication.pages,
		n.publication.publisher,
	)
}

func createMagazine(name, publisher string, pages int) *magazine {
	return &magazine{
		publication: publication{
			name:      name,
			publisher: publisher,
			pages:     pages,
		},
	}
}
