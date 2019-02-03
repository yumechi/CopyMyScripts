package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/vjeantet/jodaTime"
)

const templateDir string = "/templates/comp_template"

func currentTime() string {
	return jodaTime.Format("MMdd_HHmmss", time.Now())
}

func stringListToString(arr []string) string {
	return strings.Join(arr, "")
}

func getRealPath(path string) string {
	info, err := os.Lstat(path)
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		realPath, err := os.Readlink(path)
		if err != nil {
			panic(err)
		}
		return realPath
	}
	return path
}

func getTemplatePath(path string) string {
	_homedir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	templatePath := getRealPath(_homedir + path)
	return templatePath
}

func getCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}

func copyTemplate(from string, to string) {
	fromFile, err := os.Open(from)
	if err != nil {
		panic(err)
	}
	defer fromFile.Close()

	toFile, err := os.Create(to)
	if err != nil {
		panic(err)
	}
	defer toFile.Close()

	_, err = io.Copy(toFile, fromFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Copy: %s -> %s", from, to)
}

func correctExtension(filename string, extention string) string {
	if strings.Contains(filename, ".") {
		return filename
	}
	return filename + extention
}

func main() {
	var (
		prefix   = flag.String("prefix", "CPP", "dir and file prefix name")
		suffix   = flag.String("suffix", ".cpp", "file suffix")
		fromFile = flag.String("input", "template", "template type")
		toFile   = flag.String("output", currentTime(), "template type")
	)
	flag.Parse()

	// debug
	// fmt.Println(*prefix, *suffix, *fromFile, *toFile)

	_fromFile := correctExtension(*fromFile, *suffix)
	_toFile := correctExtension(*toFile, *suffix)
	fromFilename := stringListToString([]string{getTemplatePath(templateDir), *prefix, "/", _fromFile})
	toFilename := stringListToString([]string{getCurrentPath(), "/", _toFile})

	copyTemplate(fromFilename, toFilename)
}
