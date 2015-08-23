package main

import (
	"bytes"
	"fmt"
	"github.com/mizoguche/gitinit/fileutil"
	"log"
	"os"
	"os/exec"
	"os/user"
)

func main() {
	cmd := exec.Command("git", "init")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	homeDir := currentUser.HomeDir
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fileutil.CopyAllHooks(fmt.Sprintf("%s/.gitinit/%s/hooks", homeDir, os.Args[1]), fmt.Sprintf("%s/.git/hooks", workingDir))
}
