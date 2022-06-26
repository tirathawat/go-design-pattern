package television

type Samsung struct {
	channel int
	volume  int
	isOn    bool
}

func (tv *Samsung) setChannel(channel int) {
	tv.channel = channel
}

func (tv *Samsung) getChannel() int {
	return tv.channel
}

func (tv *Samsung) setVolume(volume int) {
	tv.volume = volume
}

func (tv *Samsung) getVolume() int {
	return tv.volume
}

func (tv *Samsung) setIsOn(isOn bool) {
	tv.isOn = isOn
}

func (tv *Samsung) getIsOn() bool {
	return tv.isOn
}

func NewSamsung(channel int, volume int) *Samsung {
	return &Samsung{
		channel: channel,
		volume:  volume,
		isOn:    false,
	}
}
