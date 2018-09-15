package report

var (
	// Error checks if the fed error
	// variable is not nil and if so is
	// followed by a message to signify
	// that something went wrong.
	Error = func(v ...interface{}) {
		for _, e := range v {
			switch e.(type) {
			case error:
				if e != nil {
					errorPrint(v...)
				}
			}
		}
	}

	// Success displays a brightly colored
	// message to signify in the logs that
	// a check or function passed.
	Success = func(v ...interface{}) {
		successPrint(v...)
	}

	// Status displays a gray colored
	// message that is used to log what the
	// application is currently working on.
	Status = func(v ...interface{}) {
		statusPrint(v...)
	}

	// Seperator will print a terminal
	// seperator to organize logs.
	Seperator = func(title string) {
		if title == "" {
			seperatorPrint()
		} else {
			seperatorPrintTitle(title)
		}
	}
)
