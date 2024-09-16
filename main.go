package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
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
	from := args[0 : len(args)-1]
	to := args[len(args)-1]
	if args[len(args)-2] == ":plus" {
		from = from[:len(from)-1]
		to = args[len(args)-2]
	}
	var newName string
	count := 0
	if len(from) != 1 && to != ":upper" && to != ":lower" && to != ":camel" && to != ":snake" && to != ":kebab" && to != ":plus" {
		fmt.Printf("Attention: you are renaming %d files.\n", len(from))
		fmt.Printf("Some files may be merged into one.\n")
		fmt.Printf("Still continue?[y/n]:")
		input := ""
		fmt.Scanln(&input)
		if input == "n" {
			fmt.Println("Rename cancelled")
			return
		}
	}
	for _, f := range from {
		if to[0] == '.' {
			newName = DoRenameExt(f, to)
		} else if to[0] == ':' {
			newName = DoTransform(f, to, args[len(args)-1])
		} else {
			newName = DoRename(f, to)
		}
		// check if newName already exists
		if fileExists(newName) && to != ":upper" && to != ":lower" {
			fmt.Printf("Error: Target file '%s' already exists\n%d files have been renamed", newName, count)
			return
		}

		//check if name is already changed
		if err := os.Rename(f, newName); err != nil {
			fmt.Println("Error:", err)
			os.Exit(2)
			return
		}
		count++
	}

	fmt.Printf("Rename success %d files\n", count)

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

// do: rename a.mp3 :upper
// will rename a.mp3 to A.mp3

// do: rename A.mp3 :lower
// will rename A.mp3 to a.mp3

// do: rename abc_def :camel
// will rename abc_def.mp3 to AbcDef.mp3

// do: rename AbcDef :snake
// will rename AbcDef.mp3 to abc_def.mp3

// do: rename abc_def :kebab
// will rename abc_def.mp3 to abc-def.mp3

// do: rename abc_def.mp3 :plus _g
// will rename abc_def.mp3 to abc_def_g.mp3

func DoTransform(from, to string, plus string) string {
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
	case ":plus":
		filename = filename + plus

	default:
		fmt.Println("Unknown transform:", to)
		os.Exit(3)
	}

	newname := dir + "/" + filename + ext
	return newname

}

// do: rename /tmp/a.mp3 b
// will rename /tmp/a.mp3 to /tmp/b.mp3
func DoRename(from, to string) string {
	dir := path.Dir(from)
	basename := path.Base(from)
	ext := path.Ext(basename)
	newname := dir + "/" + to + ext
	return newname
}

func fileExists(path string) bool {
	_, err := os.Lstat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		// 检查目录中是否存在同名文件（忽略大小写）
		dir := filepath.Dir(path)
		base := filepath.Base(path)
		entries, err := os.ReadDir(dir)
		if err != nil {
			return false
		}
		for _, entry := range entries {
			if strings.EqualFold(entry.Name(), base) {
				return true
			}
		}
	}
	return false
}
