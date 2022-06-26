package cafe

import (
	"design-pattern/structural/facade/cafe/coffeemachine"
	"fmt"
)

func MakeAmericano(size float32) {
	fmt.Printf("\nMaking an Americano\n")

	coffeeMachine := coffeemachine.New()

	beansAmount := (size / 8.0) * 5.0

	coffeeMachine.Start(beansAmount, 2)
	coffeeMachine.GrindBeans()
	coffeeMachine.UseHotWater(size)
	coffeeMachine.End()

	fmt.Println("Americano is ready")
}

func MakeLatte(size float32, foam bool) {
	fmt.Printf("\nMaking an Latte\n")

	coffeeMachine := coffeemachine.New()

	beansAmount := (size / 8.0) * 5.0
	milk := (size / 8.0) * 2.0

	coffeeMachine.Start(beansAmount, 4)
	coffeeMachine.GrindBeans()
	coffeeMachine.UseHotWater(size)
	coffeeMachine.UseMilk(milk)
	coffeeMachine.DoFoam()
	coffeeMachine.End()

	fmt.Println("Latte is ready")
}
