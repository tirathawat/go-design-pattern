package television

type Sony struct {
	channel int
	volume  int
	isOn    bool
}

func (tv *Sony) turnOn() bool {
	tv.isOn = true
	return tv.isOn
}

func (tv *Sony) turnOff() bool {
	tv.isOn = false
	return tv.isOn
}

func (tv *Sony) volumeUp() int {
	tv.volume++
	return tv.volume
}

func (tv *Sony) volumeDown() int {
	tv.volume--
	return tv.volume
}

func (tv *Sony) channelUp() int {
	tv.channel++
	return tv.channel
}

func (tv *Sony) channelDown() int {
	tv.channel--
	return tv.channel
}

func NewSony(channel int, volume int) *Sony {
	return &Sony{
		channel: channel,
		volume:  volume,
		isOn:    false,
	}
}
