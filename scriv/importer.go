// Package scriv implements importing Scrivener (.scriv) projects into Voluta.
package scriv

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/calmecac-dev/voluta/ast"
	"github.com/calmecac-dev/voluta/reader/rtf"
)

// Project represents an imported Scrivener project.
type Project struct {
	Title     string
	Documents []Document
}

// Document represents a single node in the Scrivener binder.
type Document struct {
	UUID             string
	Title            string
	Type             string // Text, Folder, DraftFolder, ResearchFolder, TrashFolder
	IncludeInCompile bool
	Doc              ast.Document
	Children         []Document
}

// --- XML structs matching Scrivener 3 .scrivx format ---

type scrivxRoot struct {
	XMLName xml.Name    `xml:"ScrivenerProject"`
	Binder  scrivBinder `xml:"Binder"`
}

type scrivBinder struct {
	Items []scrivItem `xml:"BinderItem"`
}

type scrivItem struct {
	UUID     string      `xml:"UUID,attr"`
	Type     string      `xml:"Type,attr"`
	Title    string      `xml:"Title"`
	MetaData scrivMeta   `xml:"MetaData"`
	Children []scrivItem `xml:"Children>BinderItem"`
}

type scrivMeta struct {
	IncludeInCompile string `xml:"IncludeInCompile"`
}

// Import reads a .scriv project and returns a Project.
func Import(scrivPath string, imageHandler func([]byte, string) (string, error)) (Project, error) {
	scrivxPath, err := findScrivx(scrivPath)
	if err != nil {
		return Project{}, err
	}
	data, err := os.ReadFile(scrivxPath)
	if err != nil {
		return Project{}, fmt.Errorf("scriv: cannot read %s: %w", scrivxPath, err)
	}
	var root scrivxRoot
	if err := xml.Unmarshal(data, &root); err != nil {
		return Project{}, fmt.Errorf("scriv: invalid XML: %w", err)
	}
	projectName := strings.TrimSuffix(filepath.Base(scrivxPath), ".scrivx")
	dataPath := filepath.Join(scrivPath, "Files", "Data")
	project := Project{Title: projectName}
	for _, item := range root.Binder.Items {
		if item.Type == "TrashFolder" {
			continue
		}
		doc, err := importItem(item, dataPath, imageHandler)
		if err != nil {
			return Project{}, err
		}
		project.Documents = append(project.Documents, doc)
	}
	return project, nil
}

func importItem(item scrivItem, dataPath string, imageHandler func([]byte, string) (string, error)) (Document, error) {
	doc := Document{
		UUID:             item.UUID,
		Title:            item.Title,
		Type:             item.Type,
		IncludeInCompile: item.MetaData.IncludeInCompile != "No",
	}
	rtfPath := filepath.Join(dataPath, item.UUID, "content.rtf")
	if data, err := os.ReadFile(rtfPath); err == nil {
		astDoc, err := rtf.ReadWithOptions(data, rtf.Options{
			ImageHandler: imageHandler,
		})
		if err != nil {
			return Document{}, fmt.Errorf("scriv: error reading RTF for %s: %w", item.UUID, err)
		}
		astDoc = cleanDocument(astDoc)
		astDoc.Meta.Title = item.Title
		doc.Doc = astDoc
	}
	for _, child := range item.Children {
		childDoc, err := importItem(child, dataPath, imageHandler)
		if err != nil {
			return Document{}, err
		}
		doc.Children = append(doc.Children, childDoc)
	}
	return doc, nil
}

func findScrivx(scrivPath string) (string, error) {
	entries, err := os.ReadDir(scrivPath)
	if err != nil {
		return "", fmt.Errorf("scriv: cannot open %s: %w", scrivPath, err)
	}
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".scrivx") {
			return filepath.Join(scrivPath, e.Name()), nil
		}
	}
	return "", fmt.Errorf("scriv: no .scrivx file found in %s", scrivPath)
}
