# Voluta

> A volute is a spiral ornament, resembling a snail shell, characteristic of Ionic and Composite capitals in classical architecture. In Mesoamerican codices, a similar spiral glyph — the *cuicatl* — represented the spoken word and human discourse. Voluta takes its name from both.

Voluta is a Go library for reading and converting creative writing document formats. It provides a clean, high-level API to read RTF documents and Scrivener projects, converting them to Markdown, HTML, and other formats through an intermediate AST.

Voluta is the backbone of [Tlahtolli](https://github.com/tlahtolli-app/tlahtolli), a free and open source creative writing editor, but is completely independent and available for any Go project.

## Features

- **RTF reader** — converts RTF documents to an intermediate AST, preserving bold, italic, underline, strikethrough, headings, and lists with their original markers
- **Scrivener importer** — reads `.scriv` projects (Scrivener 3), preserving the full binder tree, document metadata, and IncludeInCompile flags
- **Markdown writer** — converts the AST to clean Markdown
- **HTML writer** *(experimental)* — converts the AST to HTML
- **High-level API** — single import, simple function calls

## Installation

```bash
go get github.com/calmecac-dev/voluta
```

Requires Go 1.23 or later.

## Quick Start

### Convert RTF to Markdown

```go
import (
    "fmt"
    "log"
    "os"

    "github.com/calmecac-dev/voluta"
)

data, _ := os.ReadFile("document.rtf")

md, err := voluta.Convert(voluta.FormatRTF, voluta.FormatMarkdown, data)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(md))
```

### Read RTF into AST

```go
data, _ := os.ReadFile("document.rtf")

doc, err := voluta.Read(voluta.FormatRTF, data)
if err != nil {
    log.Fatal(err)
}

// Write to multiple formats from the same AST
md, _   := voluta.Write(voluta.FormatMarkdown, doc)
html, _ := voluta.Write(voluta.FormatHTML, doc)
```

### Import a Scrivener project

```go
project, err := voluta.ImportScriv("MyNovel.scriv")
if err != nil {
    log.Fatal(err)
}

fmt.Println("Project:", project.Title)

for _, doc := range project.Documents {
    if !doc.IncludeInCompile {
        continue
    }
    md, err := voluta.Write(voluta.FormatMarkdown, doc.Doc)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("## %s\n\n%s\n\n", doc.Title, string(md))
}
```

### Save images from RTF

```go
data, _ := os.ReadFile("document.rtf")

// Simple: provide a directory
doc, err := voluta.Read(voluta.FormatRTF, data, voluta.ReadOptions{
    ImageDir: "assets/images",
})

// Advanced: full control over image storage
doc, err := voluta.Read(voluta.FormatRTF, data, voluta.ReadOptions{
    ImageHandler: func(imgData []byte, format string) (string, error) {
        path := fmt.Sprintf("assets/%s.%s", uuid.New(), format)
        return path, os.WriteFile(path, imgData, 0644)
    },
})
```

## Supported Formats

| Format | Read | Write |
|--------|------|-------|
| RTF | Yes | — |
| Markdown | — | Yes |
| HTML | — | Experimental |
| PDF | — | Planned |
| docx | — | Planned |
| ePub | — | Planned |
| Scrivener (.scriv) | Import | Export planned |

## API Reference

### Reading

```go
// Read converts raw bytes to an ast.Document.
func Read(format Format, data []byte, opts ...ReadOptions) (ast.Document, error)
```

### Writing

```go
// Write converts an ast.Document to the target format.
func Write(format Format, doc ast.Document, opts ...WriteOptions) ([]byte, error)
```

### Converting

```go
// Convert reads from one format and writes to another in a single step.
func Convert(from, to Format, data []byte, opts ...ReadOptions) ([]byte, error)
```

### Scrivener

```go
// ImportScriv imports a Scrivener .scriv project.
func ImportScriv(path string, opts ...ReadOptions) (ScrivProject, error)
```

### Options

```go
type ReadOptions struct {
    // ImageDir is the directory where images will be saved.
    // If empty, images are skipped.
    ImageDir string

    // ImageHandler provides full control over image saving.
    // Receives image bytes and format ("png", "jpg", "wmf").
    // Should return the path or URL to use in the document.
    // If defined, takes precedence over ImageDir.
    ImageHandler func(data []byte, format string) (string, error)
}
```

## Architecture

Voluta uses a two-phase pipeline:

```
Input (RTF, .scriv)
       |
   Reader / Importer
       |
   ast.Document        <- intermediate representation
       |
   Writer
       |
Output (Markdown, HTML, ePub, PDF, docx)
```

The `ast.Document` is the central contract between readers and writers. Adding a new input or output format only requires implementing a new reader or writer — neither side needs to know about the other.

## Contributing

Voluta is open source and welcomes contributions. Please open an issue before submitting a pull request for significant changes.

## License

MIT — Copyright (c) 2026 Calmecac

---

*Part of the [Calmecac](https://calmecac.dev) open source ecosystem.*