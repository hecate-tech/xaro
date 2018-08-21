package common

import "log"

var (
	// ErrorCheck is a wrapper for error checking
	ErrorCheck = func(err error) {
		if err != nil {
			log.Fatalln(err)
		}
	}
)
