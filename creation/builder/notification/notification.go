package notification

import "fmt"

type Notification struct {
	title    string
	subtitle string
	message  string
	image    string
	icon     string
	priority int
}

func (n Notification) String() string {
	return fmt.Sprintf("Title: %s\nSubtitle: %s\nMessage: %s\nImage: %s\nIcon: %s\nPriority: %d\n",
		n.title,
		n.subtitle,
		n.message,
		n.image,
		n.icon,
		n.priority,
	)
}

func (n Notification) Validate() error {
	if n.title == "" {
		return ErrRequiredTitle
	}

	return nil
}
