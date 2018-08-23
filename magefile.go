// +build mage

package main

import (
	"fmt"
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
	mg.Deps(InstallDeps)

	fmt.Println("► Building…")

	cmd := exec.Command("go", "build", "-o", "Xaro.exe", ".")
	if err := cmd.Run(); err != nil {
		return err
	}

	log.Println("building complete")

	return nil
}

// Install will create a bin folder and place the executable in it.
func Install() error {
	mg.Deps(Build)
	fmt.Println("► Installing…")

	os.MkdirAll("./bin", os.ModePerm)
	log.Println("Created 'bin' folder")
	if err := os.Rename("./Xaro.exe", "./bin/Xaro.exe"); err != nil {
		return err
	}

	log.Println("moved executable to bin folder")

	return nil
}

// InstallDeps runs dep package manager.
func InstallDeps() error {
	fmt.Println("► Installing Deps…")

	cmd := exec.Command("dep", "ensure")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("► Cleaning…")

	os.RemoveAll("bin/Xaro.exe")
	log.Println("removed bin/Xaro.exe")
	os.RemoveAll("vendor")
	log.Println("removed vendor")
	os.RemoveAll(".vendor-new")
	log.Println("removed .vendor-new")
}
