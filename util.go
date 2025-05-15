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

// fungsi untk membaca enter

// fungsi untuk menerima input string

// fungsi untuk menerima input integer

// fungsi konversi ke lowercase

// fungsi konversi ke uppercase

// fungsi generate ID assessment

// fungsi generate tanggal

// fungsi generate ID pengguna

// fungsi mengecek ketersediaan ID pengguna