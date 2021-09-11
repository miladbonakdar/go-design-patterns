package main

import (
	"fmt"
	"time"
)

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
)

type log struct {
	level   Level
	message string
	details interface{}
	at      time.Time
}

type LogBuilder struct {
	log log
}

func newLogBuilder() *LogBuilder {
	return &LogBuilder{log: log{
		level:   Debug,
		message: "",
		details: nil,
		at:      time.Now(),
	}}
}

func (b *LogBuilder) Level(level Level) *LogBuilder {
	b.log.level = level
	return b
}

func (b *LogBuilder) Message(message string) *LogBuilder {
	b.log.message = message
	return b
}

func (b *LogBuilder) Details(data interface{}) *LogBuilder {
	b.log.details = data
	return b
}

func (b *LogBuilder) build() log {
	return b.log
}

type LogAction func(builder *LogBuilder)

func printLog(log log) {
	fmt.Println(log)
}

func Log(action LogAction) {
	builder := newLogBuilder()
	action(builder)
	log := builder.build()
	printLog(log)
}

func main() {
	Log(func(builder *LogBuilder) {
		builder.Level(Error).Message("Error has been captured").
			Details(map[string]interface{}{
				"Url": "/some-url",
			})
	})
}
