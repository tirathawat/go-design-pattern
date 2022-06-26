package main

import (
	"design-pattern/structural/adapter/television"
	"fmt"
)

func main() {
	samsungTV := television.NewSamsungAdapter(television.NewSamsung(10, 40))
	{
		fmt.Println("samsungTV")
		samsungTV.TurnOn()
		fmt.Println("Channel:", samsungTV.ChannelDown())
		samsungTV.ToChannel(100)
		fmt.Println("Channel:", samsungTV.ChannelUp())
		fmt.Println("Volume:", samsungTV.VolumeUp())
		fmt.Println("Volume:", samsungTV.VolumeDown())
		samsungTV.TurnOff()
	}

	fmt.Println()

	sonyTV := television.NewSonyAdapter(television.NewSony(20, 20))
	{
		fmt.Println("sonyTV")
		sonyTV.TurnOn()
		fmt.Println("Channel:", sonyTV.ChannelDown())
		sonyTV.ToChannel(100)
		fmt.Println("Channel:", sonyTV.ChannelUp())
		fmt.Println("Volume:", sonyTV.VolumeUp())
		fmt.Println("Volume:", sonyTV.VolumeDown())
		sonyTV.TurnOff()
	}

}
