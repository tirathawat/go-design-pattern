package coffeemachine

import "fmt"

type coffeeMachine struct {
	beanAmount   float32
	grinderLevel int
	waterTemp    int
	waterAmount  float32
	milkAmount   float32
	addFoam      bool
}

func (m *coffeeMachine) Start(beanAmount float32, grinderLevel int) {
	m.beanAmount = beanAmount
	m.grinderLevel = grinderLevel

	fmt.Println("Starting coffee order with beans:", beanAmount, "and grinder level:", grinderLevel)
}

func (m *coffeeMachine) End() {
	fmt.Println("Ending coffee order")
}

func (m *coffeeMachine) GrindBeans() bool {
	fmt.Println("Grinding the beans:", m.beanAmount, "beans at", m.grinderLevel)
	return true
}

func (m *coffeeMachine) UseMilk(amount float32) float32 {
	fmt.Println("Adding milk:", amount, "oz")
	m.milkAmount = amount
	return m.milkAmount
}

func (m *coffeeMachine) SetWaterTemp(temp int) {
	fmt.Println("Setting water temp:", temp)
	m.waterTemp = temp
}

func (m *coffeeMachine) UseHotWater(amount float32) float32 {
	fmt.Println("Adding hot water", amount)
	m.waterAmount = amount
	return m.waterAmount
}

func (m *coffeeMachine) DoFoam() {
	fmt.Println("Adding Foam")
	m.addFoam = true
}

func New() *coffeeMachine {
	return &coffeeMachine{}
}
