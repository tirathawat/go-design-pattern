package main

import "design-pattern/creation/singleton/logger"

func main() {
	log := logger.GetInstance()

	log.SetLevel(logger.Error)
	log.Log("This is log message")

	log.SetLevel(logger.Info)
	log.Log("This is log message")

	log.SetLevel(logger.Warning)
	log.Log("This is log message")

	// Use sync.Once to make singleton concurrent-safe
	for i := 0; i < 10; i++ {
		go logger.GetInstance()
	}
}
