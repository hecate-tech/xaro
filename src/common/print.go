package common

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// var delimiter = "►"

const (
	delimiter = "►"
)

var (
	statusCol  = color.CyanString
	successCol = color.HiGreenString
	failureCol = color.HiRedString
)

type colStringer func(string, ...interface{}) string

// StatusPrint will print a colored
// delimiter followed by a message.
func StatusPrint(a ...interface{}) {
	printDelimiter(statusCol)
	fmt.Print(a...)
	fmt.Println("…")
}

// errorPrint is used for ErrorCheck
// when printing the final error
func errorPrint(a ...interface{}) {
	printDelimiter(failureCol)
	color.New(color.FgHiRed).Fprintln(color.Error, a...)
	os.Exit(1)
}

// SuccessPrint is for when you want
// to display a task was completed
// successfully.
func SuccessPrint(a ...interface{}) {
	printDelimiter(successCol)
	c := color.New(color.FgHiGreen)
	c.Fprint(color.Output, a...)
	c.Fprint(color.Output, "…\n")
}

func printDelimiter(col colStringer) {
	fmt.Fprintf(color.Output, col(delimiter+" "))
}
