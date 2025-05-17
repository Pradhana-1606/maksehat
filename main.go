package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Contoh Penggunaan: go run main.go [cli/gui]")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "cli":
		cliMode()
	case "gui":
		guiMode()
	default:
		fmt.Println("Perintah tidak dikenali:", mode)
		fmt.Println("Contoh Penggunaan: go run main.go cli (untuk masuk ke mode CLI)")
	}
}