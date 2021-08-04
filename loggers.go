package loggers

import (
	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	invalidPhoneMessage     = Event{"Invalid phone number found"}
	invalidNameValueMessage = Event{"Empty Name found"}
	invalidEmailMessage     = Event{"Invalid Email found"}
	invalidIDMessage        = Event{"Duplicate User ID found..Assigning new User ID"}
	errorMessage            = Event{"Error:%s"}
)

func (l *StandardLogger) InvalidPhone() {
	l.Errorf(invalidPhoneMessage.message)
}
func (l *StandardLogger) InvalidName() {
	l.Errorf(invalidNameValueMessage.message)
}
func (l *StandardLogger) InvalidEmail() {
	l.Errorf(invalidEmailMessage.message)
}
func (l *StandardLogger) DuplicateEntry() {
	l.Errorf(invalidIDMessage.message)
}
func (l *StandardLogger) Error(arg string) {
	l.Errorf(errorMessage.message, arg)
}
