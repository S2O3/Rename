package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.Args)
	// fmt.Println("len(os.Args):", len(os.Args))
	// fmt.Println(os.Args[1:])
	DoCommand(os.Args[1:])
}

func DoCommand(args []string) {
	if len(args) == 0 || args[0] == "-h" {
		ShowHelp()
		return
	}

}

func ShowHelp() {
	fmt.Println("$ rename [from] [to] [options...]")
}
