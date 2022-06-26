package television

type samsungAdapter struct {
	tv *Samsung
}

func (a *samsungAdapter) VolumeUp() int {
	a.tv.setVolume(a.tv.volume + 1)
	return a.tv.getVolume()
}

func (a *samsungAdapter) VolumeDown() int {
	a.tv.setVolume(a.tv.volume - 1)
	return a.tv.getVolume()
}

func (a *samsungAdapter) ChannelUp() int {
	a.tv.setChannel(a.tv.channel + 1)
	return a.tv.getChannel()
}

func (a *samsungAdapter) ChannelDown() int {
	a.tv.setChannel(a.tv.channel - 1)
	return a.tv.getChannel()
}

func (a *samsungAdapter) TurnOn() {
	a.tv.setIsOn(true)
}

func (a *samsungAdapter) TurnOff() {
	a.tv.setIsOn(false)
}

func (a *samsungAdapter) ToChannel(channel int) {
	a.tv.setChannel(channel)
}

func NewSamsungAdapter(tv *Samsung) *samsungAdapter {
	return &samsungAdapter{tv: tv}
}
