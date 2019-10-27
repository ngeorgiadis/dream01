package intlog

import (
	"fmt"
	"log"
	"time"

	"dream01/internal/lumberjack"
)

// IntLogger ...
type IntLogger struct {
	l *lumberjack.Logger
}

// NewIntLogger ...
func NewIntLogger(l *lumberjack.Logger) *IntLogger {
	return &IntLogger{
		l: l,
	}
}

// Write ...
func (ml *IntLogger) Write(p []byte) (n int, err error) {

	fmt.Println(string(p))

	pre := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	p = append([]byte(pre+" "), p...)
	return ml.l.Write(p)
}

// Info ...
func Info(s string) {
	log.Printf("[%s] %s", "INFO", s)
}

// Infof ...
func Infof(f string, v ...interface{}) {
	s := fmt.Sprintf(f, v...)
	Info(s)
}

// Warn ...
func Warn(s string) {
	log.Printf("[%s] %s", "WARN", s)
}

// Warnf ...
func Warnf(f string, v ...interface{}) {
	s := fmt.Sprintf(f, v...)
	Warn(s)
}

// Error ...
func Error(s string) {
	log.Printf("[%s] %s", "ERROR", s)
}

// Errorf ...
func Errorf(f string, v ...interface{}) {
	s := fmt.Sprintf(f, v...)
	Error(s)
}
