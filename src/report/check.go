package report

var (
	// Error is a wrapper for error checking
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
)
