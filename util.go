package main

import (
	"bufio"
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

// fungsi untk membaca enter
func pressEnter() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

// fungsi konversi ke lowercase

// fungsi konversi ke uppercase

// fungsi generate ID assessment

// fungsi generate tanggal

// fungsi generate ID pengguna

// fungsi mengecek ketersediaan ID pengguna