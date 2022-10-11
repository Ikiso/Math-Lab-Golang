package matrix

import "fmt"

var DEBUG bool = false

func SetDebug(debug bool) {
	DEBUG = debug
}

func GenerateError(message string) (err error) {
	if DEBUG {
		panic(message)
	} else {
		err = fmt.Errorf(message)
	}

	return
}
