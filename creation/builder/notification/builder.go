package notification

type Builder struct {
	notification Notification
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (nb *Builder) SetTitle(title string) *Builder {
	nb.notification.title = title
	return nb
}

func (nb *Builder) SetSubtitle(subtitle string) *Builder {
	nb.notification.subtitle = subtitle
	return nb
}

func (nb *Builder) SetMessage(message string) *Builder {
	nb.notification.message = message
	return nb
}

func (nb *Builder) SetImage(image string) *Builder {
	nb.notification.image = image
	return nb
}

func (nb *Builder) SetIcon(icon string) *Builder {
	nb.notification.icon = icon
	return nb
}

func (nb *Builder) SetPriority(priority int) *Builder {
	nb.notification.priority = priority
	return nb
}

func (nb *Builder) Build() (*Notification, error) {
	if err := nb.notification.Validate(); err != nil {
		return nil, err
	}

	return &nb.notification, nil
}
