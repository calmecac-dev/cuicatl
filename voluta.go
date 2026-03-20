// Package voluta provides a high-level API for reading and converting
// creative writing document formats.
package voluta

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/calmecac-dev/voluta/ast"
	rtfreader "github.com/calmecac-dev/voluta/reader/rtf"
	htmlwriter "github.com/calmecac-dev/voluta/writer/html"
	markdownwriter "github.com/calmecac-dev/voluta/writer/markdown"
)

// Format identifies a document format.
type Format string

const (
	FormatRTF      Format = "rtf"
	FormatMarkdown Format = "markdown"
	FormatHTML     Format = "html"
	FormatPDF      Format = "pdf"
	FormatDocx     Format = "docx"
	FormatEPub     Format = "epub"
)

// ReadOptions configures the Read function.
type ReadOptions struct {
	// ImageDir is the directory where images will be saved.
	// If empty, images are skipped.
	ImageDir string

	// ImageHandler provides full control over image saving.
	// If defined, takes precedence over ImageDir.
	ImageHandler func(data []byte, format string) (string, error)
}

// Read converts raw document data to an ast.Document.
func Read(format Format, data []byte, opts ...ReadOptions) (ast.Document, error) {
	var o ReadOptions
	if len(opts) > 0 {
		o = opts[0]
	}

	handler := o.ImageHandler
	if handler == nil && o.ImageDir != "" {
		handler = defaultImageHandler(o.ImageDir)
	}

	switch format {
	case FormatRTF:
		return rtfreader.ReadWithOptions(data, rtfreader.Options{
			ImageHandler: handler,
		})
	default:
		return ast.Document{}, fmt.Errorf("voluta: unsupported read format %q", format)
	}
}

// WriteOptions configures the Write function.
type WriteOptions struct{}

// Write converts an ast.Document to the target format.
func Write(format Format, doc ast.Document, opts ...WriteOptions) ([]byte, error) {
	switch format {
	case FormatMarkdown:
		out, err := markdownwriter.Write(doc)
		return []byte(out), err
	case FormatHTML:
		out, err := htmlwriter.Write(doc)
		return []byte(out), err
	default:
		return nil, fmt.Errorf("voluta: unsupported write format %q", format)
	}
}

// Convert reads data from one format and writes it to another in a single step.
func Convert(from, to Format, data []byte, opts ...ReadOptions) ([]byte, error) {
	doc, err := Read(from, data, opts...)
	if err != nil {
		return nil, err
	}
	return Write(to, doc)
}

func defaultImageHandler(dir string) func([]byte, string) (string, error) {
	return func(data []byte, format string) (string, error) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", fmt.Errorf("voluta: cannot create image dir: %w", err)
		}
		filename := fmt.Sprintf("img_%d.%s", time.Now().UnixNano(), format)
		path := filepath.Join(dir, filename)
		return path, os.WriteFile(path, data, 0644)
	}
}
