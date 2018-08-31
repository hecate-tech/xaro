// +build mage

package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Install

// Build will clean, install dependencies, and compile the project into an executable
func Build() error {
	mg.Deps(Clean)       // Clean old executables and vendor folder/s
	mg.Deps(RunFmt)      // Runs gofmt to make sure everything is formatted correctly.
	mg.Deps(BuildDep)    // Updates and builds dep.exe in your gobin folder
	mg.Deps(InstallDeps) // Uses dep.exe to install the dependencies for engo-xaro
	mg.Deps(RunLinter)   // Runs golint to find bad practices in the source code.

	fmt.Println("► Building Executables…")

	// go build -o Xaro.exe ./src/cmd/.
	cmd := exec.Command("go", "build", "-o", "Xaro.exe", "./src/cmd/.")
	if err := cmd.Run(); err != nil {
		return err
	}

	log.Println("built Xaro.exe…")

	// go build -o Server.exe ./src/communication/server/.
	cmd = exec.Command("go", "build", "-o", "Server.exe", "./src/communication/server/.")
	if err := cmd.Run(); err != nil {
		return err
	}

	log.Println("built Server.exe…")

	return nil
}

// Install will create a bin folder and place the executable in it.
func Install() error {
	mg.Deps(Build)
	fmt.Println("► Installing…")

	// Creates the bin directory in the root directory.
	os.MkdirAll("./bin", os.ModePerm)
	log.Println("created 'bin' folder…")

	// Moves Xaro.exe into the new bin folder
	if err := os.Rename("./Xaro.exe", "bin/Xaro.exe"); err != nil {
		return err
	}
	log.Println("moved Xaro.exe to bin…")

	// Moves Server.exe into the new bin folder
	if err := os.Rename("./Server.exe", "bin/Server.exe"); err != nil {
		return err
	}
	log.Println("moved Server.exe to bin…")

	log.Println("moved executable to bin folder…")

	return nil
}

// RunFmt runs gofmt to format the src files
func RunFmt() error {
	fmt.Println("► Running gofmt…")
	// gofmt -l -w .
	cmd := exec.Command("gofmt", "-l", "-w", ".")

	return cmd.Run()
}

func BuildDep() error {
	gopath := build.Default.GOPATH

	// checks if dep exists. If it doesn't then return the error.
	if _, err := os.Stat(gopath + "/src/github.com/golang/dep"); os.IsNotExist(err) {
		return err
	}

	fmt.Println("► Building github.com/golang/dep/...…")

	// go get -u github.com/golang/dep/cmd/dep
	cmd := exec.Command("go", "get", "-u", "github.com/golang/dep/cmd/dep")
	// builds and install dep into the bin folder.
	return cmd.Run()
}

// InstallDeps runs dep package manager.
func InstallDeps() error {
	fmt.Println("► Installing Dependencies…")

	// ensures that you have vendored files in order
	// for the project to run successfully.
	cmd := exec.Command("dep", "ensure")

	return cmd.Run()
}

// RunLinter runs golint to catch bad habits...
func RunLinter() error {
	fmt.Println("► Running golint…")

	// golint -set_exit_status ./src/...
	cmd := exec.Command("golint", "-set_exit_status", "./src/...")
	b, err := cmd.Output() // Stores the cmd output to b ([]byte)
	if err != nil {        // if the command exited with status code 1
		fmt.Print(string(b)) // Prints out the error/output as string
		return err
	}

	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("► Cleaning…")

	os.Remove("bin/Xaro.exe")
	log.Println("removed bin/Xaro.exe…")

	os.Remove("bin/Server.exe")
	log.Println("removed bin/Server.exe…")

	os.RemoveAll("vendor")
	log.Println("removed vendor…")

	os.RemoveAll(".vendor-new…")
	log.Println("removed .vendor-new…")
}
