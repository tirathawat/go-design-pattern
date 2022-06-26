package television

type sonyAdapter struct {
	tv *Sony
}

func (a *sonyAdapter) VolumeUp() int {
	return a.tv.volumeUp()
}

func (a *sonyAdapter) VolumeDown() int {
	return a.tv.volumeDown()
}

func (a *sonyAdapter) ChannelUp() int {
	return a.tv.channelUp()
}

func (a *sonyAdapter) ChannelDown() int {
	return a.tv.channelDown()
}

func (a *sonyAdapter) TurnOn() {
	a.tv.turnOn()
}

func (a *sonyAdapter) TurnOff() {
	a.tv.turnOff()
}

func (a *sonyAdapter) ToChannel(channel int) {
	for channel != a.tv.channel {
		if channel < a.tv.channel {
			a.tv.channelDown()
		} else {
			a.tv.channelUp()
		}
	}
}

func NewSonyAdapter(tv *Sony) *sonyAdapter {
	return &sonyAdapter{tv: tv}
}
