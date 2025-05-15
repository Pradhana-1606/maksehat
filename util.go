package main

import (
	"os"
	"os/exec"
	"runtime"
)

// fungsi untuk membersihkan terminal
func clearConsole(){
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
