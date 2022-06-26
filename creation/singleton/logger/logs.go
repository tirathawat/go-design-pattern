package logger

import (
	"fmt"
	"sync"
)

type Level string

const (
	Debug   Level = "debug"
	Info    Level = "info"
	Warning Level = "warning"
	Error   Level = "error"

	Default Level = Debug
)

type logger struct {
	level Level
}

var instance *logger
var once sync.Once

func (l *logger) SetLevel(level Level) {
	l.level = level
}

func (l *logger) Log(msg string) {
	fmt.Println(l.level, ":", msg)
}

func newLogger(level Level) *logger {
	return &logger{
		level: level,
	}
}

func GetInstance() *logger {
	once.Do(func() {
		instance = newLogger(Default)
	})

	return instance
}
