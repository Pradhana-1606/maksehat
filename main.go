package main

import (
	"fmt"
	"maksehat/cli"
	"maksehat/gui"
	"math/rand"
	"os"
	"time"
)

func main(){
	rand.New(rand.NewSource(time.Now().UnixNano()))

	if len(os.Args) < 2 {
		fmt.Println("Contoh Penggunaan: go run main.go . [cli/gui]")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "cli":
		cli.CliMode()
	case "gui":
		gui.GuiMode()
	default:
		fmt.Println("Perintah tidak dikenali:", mode)
		fmt.Println("Contoh Penggunaan: go run main.go . cli (untuk masuk ke mode CLI)")
	}
}