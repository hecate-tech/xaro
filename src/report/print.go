package report

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

const (
	prefix = "►"
)

var (
	statusCol  = color.CyanString
	successCol = color.HiGreenString
	failureCol = color.HiRedString
)

type colStringer func(string, ...interface{}) string

// Status will print a colored
// delimiter followed by a message.
func Status(a ...interface{}) {
	printPrefix(statusCol)
	fmt.Print(a...)
	fmt.Println("…")
}

// errorPrint is used for ErrorCheck
// when printing the final error
func errorPrint(a ...interface{}) {
	printPrefix(failureCol)
	color.New(color.FgHiRed).Fprintln(color.Error, a...)
	os.Exit(1)
}

// Success is for when you want
// to display a task was completed
// successfully.
func Success(a ...interface{}) {
	printPrefix(successCol)
	c := color.New(color.FgHiGreen)
	c.Fprint(color.Output, a...)
	c.Fprint(color.Output, "…\n")
}

func printPrefix(col colStringer) {
	fmt.Fprintf(color.Output, col(prefix+" "))
}
