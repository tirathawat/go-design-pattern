package subject

type observable interface {
	Register(obs observer)
	Unregister(obs observer)
	notify()
}

type subject struct {
	observers Listeners
	field     string
}

func (s *subject) ChangeItem(data string) {
	s.field = data
	s.notify()
}

func (s *subject) Register(listener *Listener) {
	s.observers.add(listener)
}

func (s *subject) Unregister(listener *Listener) {
	s.observers.remove(listener)
}

func (s *subject) notify() {
	for _, obs := range s.observers {
		obs.onUpdate(s.field)
	}
}

func New() *subject {
	return &subject{
		observers: make(Listeners, 0),
	}
}
