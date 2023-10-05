package logging

import (
	"log"
)

type StandardLogger struct{}

func (l *StandardLogger) Debug(msg string, keysAndValues ...interface{}) {
	log.Println(append([]interface{}{"DEBUG: ", msg}, keysAndValues...)...)
}

func (l *StandardLogger) Info(msg string, keysAndValues ...interface{}) {
	log.Println(append([]interface{}{"INFO: ", msg}, keysAndValues...)...)
}

func (l *StandardLogger) Warn(msg string, keysAndValues ...interface{}) {
	log.Println(append([]interface{}{"WARN: ", msg}, keysAndValues...)...)
}

func (l *StandardLogger) Error(msg string, keysAndValues ...interface{}) {
	log.Println(append([]interface{}{"ERROR: ", msg}, keysAndValues...)...)
}

func (l *StandardLogger) Fatal(msg string, keysAndValues ...interface{}) {
	log.Fatalln(append([]interface{}{"FATAL: ", msg}, keysAndValues...)...)
}
