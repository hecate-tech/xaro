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
	mg.Deps(Clean)
	mg.Deps(RunFmt)
	mg.Deps(BuildDep)
	mg.Deps(InstallDeps)
	mg.Deps(RunLinter)

	fmt.Println("► Building Executables…")

	cmd := exec.Command("go", "build", "-o", "Xaro.exe", "./src/cmd/.")
	if err := cmd.Run(); err != nil {
		return err
	}

	log.Println("built Xaro.exe…")

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

	os.MkdirAll("./bin", os.ModePerm)
	log.Println("created 'bin' folder…")
	if err := os.Rename("./Xaro.exe", "bin/Xaro.exe"); err != nil {
		return err
	}
	log.Println("moved Xaro.exe to bin…")

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
	cmd := exec.Command("gofmt", "-l", "-w", ".")

	return cmd.Run()
}

func BuildDep() error {
	gopath := build.Default.GOPATH
	if _, err := os.Stat(gopath + "/src/github.com/golang/dep"); os.IsNotExist(err) {
		return err
	}

	fmt.Println("► Building github.com/golang/dep/...…")
	cmd := exec.Command("go", "get", "-u", "github.com/golang/dep/cmd/dep")

	return cmd.Run()
}

// InstallDeps runs dep package manager.
func InstallDeps() error {
	fmt.Println("► Installing Dependencies…")

	cmd := exec.Command("dep", "ensure")
	return cmd.Run()
}

// RunLinter runs golint to catch bad habits...
func RunLinter() error {
	fmt.Println("► Running golint…")

	cmd := exec.Command("golint", "-set_exit_status", "./src/...")
	b, err := cmd.Output()
	if err != nil {
		fmt.Print(string(b))
		return err
	}

	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("► Cleaning…")

	os.RemoveAll("bin/Xaro.exe")
	log.Println("removed bin/Xaro.exe…")

	os.RemoveAll("bin/Server.exe")
	log.Println("removed bin/Server.exe…")

	os.RemoveAll("vendor")
	log.Println("removed vendor…")

	os.RemoveAll(".vendor-new…")
	log.Println("removed .vendor-new…")
}
