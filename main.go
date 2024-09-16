package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/iancoleman/strcase"
)

func main() {
	DoCommand(os.Args[1:])
}

func DoCommand(args []string) {
	if len(args) < 2 || args[0] == "-h" {
		ShowHelp()
		return
	}
	from := args[0]
	to := args[1]
	var newName string
	if to[0] == '.' {
		newName = DoRenameExt(from, to)
	} else if to[0] == ':' {
		newName = DoTransform(from, to, args)
	} else {
		newName = DoRename(from, to)
	}
	if err := os.Rename(from, newName); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		fmt.Printf("Rename success: %s -> %s\n", from, newName)
	}

}

func ShowHelp() {
	fmt.Println("Usage:$ rename [from] [to] [options...]")
}

// do: rename /tmp/a.mp3 .txt
// will rename /tmp/a.mp3 to /tmp/a.txt
func DoRenameExt(from, to string) string {
	dir := path.Dir(from)
	basename := path.Base(from)
	//if name is music.mp3 then filename is music and ext is .mp3
	filename := strings.Split(basename, ".")[0]
	newext := to
	newname := dir + "/" + filename + newext
	return newname
}

func DoTransform(from, to string, args []string) string {
	dir := path.Dir(from)
	basename := path.Base(from)
	filename := strings.Split(basename, ".")[0]
	ext := path.Ext(basename)
	switch to {
	case ":upper":
		filename = strings.ToUpper(filename)
	case ":lower":
		filename = strings.ToLower(filename)
	case ":camel":
		filename = strcase.ToCamel(filename)
	case ":snake":
		filename = strcase.ToSnake(filename)
	case ":kebab":
		filename = strcase.ToKebab(filename)
	case ":camel_lower":
		filename = strcase.ToLowerCamel(filename)
	case ":plus":
		filename = filename + args[2]
	default:
		fmt.Println("Unknown transform:", to)
		os.Exit(3)
	}

	newname := dir + "/" + filename + ext
	return newname

}

func DoRename(from, to string) string {
	//do: rename /tmp/a.mp3 b
	//will rename /tmp/a.mp3 to /tmp/b.mp3
	dir := path.Dir(from)
	basename := path.Base(from)
	ext := path.Ext(basename)
	newname := dir + "/" + to + ext
	return newname
}
