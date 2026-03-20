package main

import (
	"fmt"
	"os"

	"github.com/calmecac-dev/voluta/reader/rtf"
	"github.com/calmecac-dev/voluta/writer/markdown"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "uso: rtf2md <archivo.rtf>")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error leyendo archivo: %v\n", err)
		os.Exit(1)
	}

	doc, err := rtf.Read(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parseando RTF: %v\n", err)
		os.Exit(1)
	}

	md, err := markdown.Write(doc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error convirtiendo a markdown: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(md)
}
