package publication

func New(
	publicationType PublicationType,
	name string,
	publisher string,
	pages int,
) (Publication, error) {

	if publicationType == Newspaper {
		return createNewspaper(name, publisher, pages), nil
	}

	if publicationType == Magazine {
		return createMagazine(name, publisher, pages), nil
	}

	return nil, ErrInvalidType
}
