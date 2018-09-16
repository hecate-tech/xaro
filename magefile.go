// +build mage

package main

import (
	"errors"
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
)

var (
	// Default target to run when none is specified
	// If not set, running mage will list available targets
	Default = Install
)

// Build will clean, install dependencies, and compile the project into an executable
func Build() error {
	mg.Deps(Clean, CheckEnv)           // Clean old executables/vendored folders and check env variables.
	mg.Deps(BuildDep)                  // Rebuild golang/dep to make sure it's up to date or installed.
	mg.Deps(InstallDeps)               // Fetches required dependencies and stores them in vendor folder/s.
	mg.Deps(RunFmt, RunLinter, RunVet) // Run all code checks for bad practices and suspicious constructs.

	fmt.Println("► Building Executables…")

	// go build -o Xaro.exe ./src/cmd/.
	cmd := exec.Command("go", "build", "-o", "Xaro.exe", "./src/cmd/client/.")
	if err := cmd.Run(); err != nil {
		return err
	}

	log.Println("built Xaro.exe…")

	return nil
}

// Install will create a bin folder and place the executable in it.
func Install() error {
	mg.Deps(Build)
	fmt.Println("► Installing…")

	// Creates the bin directory in the root directory.
	os.MkdirAll("./bin", os.ModePerm)
	log.Println("   created 'bin' folder…")

	// Moves Xaro.exe into the new bin folder
	if err := os.Rename("./Xaro.exe", "bin/Xaro.exe"); err != nil {
		return err
	}
	log.Println("Xaro successfully installed!")

	return nil
}

// BuildDep and install dep into the bin folder.
func BuildDep() error {
	fmt.Println("► Building github.com/golang/dep/...…")

	// go get -u github.com/golang/dep/cmd/dep
	cmd := exec.Command("go", "get", "-u", "github.com/golang/dep/cmd/dep")
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(b))
		return err
	}

	return nil
}

// InstallDeps runs dep package manager.
func InstallDeps() error {
	fmt.Println("► Installing Dependencies…")

	// ensures that you have vendored files in order
	// for the project to run successfully.
	cmd := exec.Command("dep", "ensure")
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(b))
		return err
	}

	return nil
}

// CheckEnv assures that the environment variables are set
func CheckEnv() error {
	log.Println("► Checking Environment Variables…")

	envVars := map[string]string{
		"gopath": build.Default.GOPATH,
		"goroot": build.Default.GOROOT,
		"goarch": build.Default.GOARCH,
		"goos":   build.Default.GOOS,
	}

	for i, v := range envVars {
		if v == "" {
			fmt.Printf("%s not found.", i)
			return errors.New("Couldn't find environment variable.")
		}
		log.Printf("   %s: %s", i, v)
	}

	return nil
}

// RunFmt runs gofmt to format the src files
func RunFmt() error {
	fmt.Println("► Running gofmt…")

	// gofmt -l -w .
	cmd := exec.Command("gofmt", "-l", "-w", ".")
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(b))
		return err
	}

	return nil
}

// RunLinter runs golint to catch bad habits...
func RunLinter() error {
	fmt.Println("► Running golint…")

	// golint -set_exit_status ./src/...
	cmd := exec.Command("golint", "-set_exit_status", "./src/...")
	b, err := cmd.Output() // Stores the cmd output to b ([]byte)
	if err != nil {        // if the command exited with an error
		fmt.Print(string(b)) // Prints out the error/output as string
		return err
	}

	return nil
}

// RunVet checks for suspicious constructs.
func RunVet() error {
	fmt.Println("► Running govet…")

	// go vet ./...
	cmd := exec.Command("go", "vet", "./src/...")

	b, err := cmd.CombinedOutput() // Stores the cmd output to b (b []byte)
	if err != nil {                // if the command exited with an error
		fmt.Print(string(b)) // prints out the result
		return err
	}

	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("► Cleaning…")

	os.Remove("bin/Xaro.exe")
	log.Println("   removed bin/Xaro.exe…")

	os.Remove("bin/Server.exe")
	log.Println("   removed bin/Server.exe…")

	os.RemoveAll("vendor")
	log.Println("   removed vendor…")

	os.RemoveAll(".vendor-new…")
	log.Println("   removed .vendor-new…")
}
