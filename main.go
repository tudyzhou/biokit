package main

import (
	"bytes"
	"fmt"
	"github.com/tudyzhou/biokit/unique"
	"os"
	"path"
	"runtime"
)

// Auther and version
const (
	Version = "1.0"
	Auther  = "tudyzhb@gmail.com"
)

// Package introduce
var (
	Packages = []packageIntron{
		{"unique", "uniuqe file"},
	}
)

type packageIntron struct {
	Name   string
	Intron string
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	// Packages introduce
	var intronBuffer bytes.Buffer
	for _, intron := range Packages {
		intronBuffer.WriteByte('\n')
		intronBuffer.WriteString(fmt.Sprintf(`			%s		%s`, intron.Name, intron.Intron))
	}

	// USAGE
	Usage := fmt.Sprintf(`
	USAGE:
		Version: %s
		Auther: %s
		A kit of bioinformatic.

		%s <package> [argv1, argv2, ...]
	
		<pacake>:%s
	`, Version, Auther, path.Base(os.Args[0]), intronBuffer.String())

	if len(os.Args) < 2 {
		fmt.Println(Usage)
	} else {
		switch t := os.Args[1]; t {
		case "unique":
			unique.Main(os.Args[2:])
		default:
			fmt.Println(Usage)
		}
	}
}
