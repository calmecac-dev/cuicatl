package main

import (
	"fmt"
	"os"

	"github.com/calmecac-dev/cuicatl/scriv"
	"github.com/calmecac-dev/cuicatl/writer/markdown"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "uso: scrivimport <proyecto.scriv>")
		os.Exit(1)
	}

	project, err := scriv.Import(os.Args[1], nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error importando proyecto: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("# %s\n\n", project.Title)
	printDocuments(project.Documents, 0)
}

func printDocuments(docs []scriv.Document, depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}

	for _, doc := range docs {
		compile := ""
		if !doc.IncludeInCompile {
			compile = " [excluido]"
		}

		fmt.Printf("%s## %s [%s]%s\n\n", indent, doc.Title, doc.Type, compile)

		if len(doc.Doc.Children) > 0 {
			md, err := markdown.Write(doc.Doc)
			if err == nil && md != "" {
				fmt.Println(md)
				fmt.Println()
			}
		}

		if len(doc.Children) > 0 {
			printDocuments(doc.Children, depth+1)
		}
	}
}
