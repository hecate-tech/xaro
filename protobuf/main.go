package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := "protoc -I="

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("error finding working directory: ", err)
	}

	cmd += wd + " --go_out=" + strings.TrimRight(wd, "protobuf/") + "src/proto/ " + wd + "/engo-xaro.proto"

	c := exec.Command("cmd", "/C", cmd)

	if err := c.Run(); err != nil {
		log.Fatalln("command failed to execute: ", err)
	}
}
