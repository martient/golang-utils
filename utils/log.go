package utils

import (
	log "github.com/sirupsen/logrus"
)

func LogInfo(message string, service string, args ...interface{}) {
	if len(service) != 0 {
		log.WithFields(log.Fields{
			"service": service,
		}).Infof(message, args...)
	} else {
		log.Infof(message, args...)
	}
}

func LogWarning(message string, service string, args ...interface{}) {
	if len(service) != 0 {
		log.WithFields(log.Fields{
			"service": service,
		}).Warnf(message, args...)
	} else {
		log.Warnf(message, args...)
	}
}

func LogError(message string, service string, args ...interface{}) {
	if len(service) != 0 {
		log.WithFields(log.Fields{
			"service": service,
		}).Errorf(message, args...)
	} else {
		log.Errorf(message, args...)
	}
}

func LogDebug(message string, service string, args ...interface{}) {
	if len(service) != 0 {
		log.WithFields(log.Fields{
			"service": service,
		}).Debugf(message, args...)
	} else {
		log.Debugf(message, args...)
	}
}
