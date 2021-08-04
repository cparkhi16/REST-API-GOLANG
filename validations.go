package main

import (
	"user/loggers"
)

func checkName(s string, standardlogger *loggers.StandardLogger) bool {
	if s == "" {
		standardlogger.InvalidName()
		return false
	}
	return true
}
func checkPhone(s string, standardlogger *loggers.StandardLogger) bool {
	if len(s) == 10 {
		return true
	} else {
		standardlogger.InvalidPhone()
		return false
	}
}
func checkEmail(email string, standardlogger *loggers.StandardLogger) bool {
	if email == "" {
		standardlogger.InvalidEmail()
		return false
	}
	return true
}
func checkValidity(name string, phonenumber string, email string, standardlogger *loggers.StandardLogger) bool {

	if checkName(name, standardlogger) && checkPhone(phonenumber, standardlogger) && checkEmail(email, standardlogger) {
		return true
	}
	return false
}
