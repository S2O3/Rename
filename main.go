package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	// fmt.Println(os.Args)
	// fmt.Println("len(os.Args):", len(os.Args))
	// fmt.Println(os.Args[1:])
	DoCommand(os.Args[1:])
}

func DoCommand(args []string) {
	if len(args) < 2 || args[0] == "-h" {
		ShowHelp()
		return
	}
	from := args[0]
	to := args[1]
	fmt.Printf("from: %s\nto: %s\n", from, to)

	if to[0] == '.' {
		DoRenameExt(from, to)
		return
	}

	//do: rename /tmp/a.mp3 b
	//will rename /tmp/a.mp3 to /tmp/b.mp3
	dir := path.Dir(from)
	basename := path.Base(from)
	ext := path.Ext(basename)
	newname := dir + "/" + to + ext
	if err := os.Rename(from, newname); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		fmt.Println("Rename success")
	}
}

func ShowHelp() {
	fmt.Println("Usage:$ rename [from] [to] [options...]")
}

// do: rename /tmp/a.mp3 .txt
// will rename /tmp/a.mp3 to /tmp/a.txt
func DoRenameExt(from, to string) {
	dir := path.Dir(from)
	basename := path.Base(from)
	filename := strings.Split(basename, ".")[0]
	newext := to
	newname := dir + "/" + filename + newext
	if err := os.Rename(from, newname); err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	} else {
		fmt.Println("Rename success")
	}
}
