package television

type Television interface {
	VolumeUp() int
	VolumeDown() int
	ChannelUp() int
	ChannelDown() int
	TurnOn()
	TurnOff()
	ToChannel(channel int)
}
