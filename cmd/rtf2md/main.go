package main

import (
	"fmt"
	"os"

	"github.com/calmecac-dev/voluta"
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

	out, err := voluta.Convert(voluta.FormatRTF, voluta.FormatMarkdown, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error convirtiendo: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
