package main

import "design-pattern/structural/facade/cafe"

func main() {
	cafe.MakeAmericano(8.0)
	cafe.MakeLatte(12.0, true)
}
