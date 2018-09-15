package report

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

const (
	prefix = "►"
	suffix = "…"
	sep    = "="

	statusCol  = color.FgCyan
	successCol = color.FgHiGreen
	failureCol = color.FgHiRed

	seperatorLen = 48
)

func errorPrint(v ...interface{}) {
	printMsg(failureCol, true, v...)
	os.Exit(1)
}

func successPrint(v ...interface{}) {
	printMsg(successCol, true, v...)
}

func statusPrint(v ...interface{}) {
	printMsg(statusCol, false, v...)
}

func seperatorPrintTitle(title string) {
	sepCount := (seperatorLen - 2) - len(title) // 2 for the brackets

	printSegment(sepCount / 2)

	fmt.Printf("[%s]", title)

	printSegment(sepCount / 2)

	fmt.Println()
}

func seperatorPrint() {
	printSegment(seperatorLen)
	fmt.Println()
}

func printSegment(amount int) {
	for i := 0; i < amount; i++ {
		fmt.Print(sep)
	}
}

func printMsg(col color.Attribute, wholemsg bool, msgs ...interface{}) {
	msgCol := color.New(col)
	msgCol.Fprint(color.Output, prefix+" ")
	if wholemsg {
		msgCol.Fprint(color.Output, msgs...)
		msgCol.Fprintln(color.Output, suffix)
		return
	}

	fmt.Print(msgs...)
	fmt.Println(suffix)
}
