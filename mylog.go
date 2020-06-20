package mylog_debugger

import (
"log"
)

type MyLog struct {
	LogLevel string
}

func (m *MyLog) Check(err error, args ...interface{}) {
	if err != nil {
		if args != nil {
			m.Print(args...)
		}
		m.Fatal("Error: ", err.Error())
	}
}

func (m *MyLog) Debug(args ...interface{}) {
	if m.LogLevel == "debug" {
		m.Print(args...)
	}
}

func (m *MyLog) Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func (m *MyLog) Print(args ...interface{}) {
	log.Print(args...)
}

func (m *MyLog) SetLogLevel(level string) {
	m.LogLevel = level
	m.Debug("Log Level: %s\n", m.LogLevel)
}