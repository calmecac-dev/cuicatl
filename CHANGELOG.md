# Changelog

All notable changes to Voluta will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0-alpha] - 2026-03-23

Initial alpha release. Core RTF reading and Scrivener importing pipeline.

### Added

**High-level API (`voluta.go`)**
- `Read(format, data, opts)` — converts raw bytes to `ast.Document`
- `Write(format, doc, opts)` — converts `ast.Document` to target format
- `Convert(from, to, data, opts)` — reads and writes in a single step
- `ImportScriv(path, opts)` — imports a Scrivener `.scriv` project
- `ReadOptions` with `ImageDir` and `ImageHandler` for image extraction

**AST (`ast/`)**
- `Document` and `Node` types as the intermediate representation
- Node types: `Paragraph`, `Heading`, `Text`, `Bold`, `Italic`, `Underline`, `Strikethrough`, `LineBreak`, `HorizontalRule`, `BlockQuote`, `Image`, `List`, `ListItem`, `Table`, `TableRow`, `TableCell`
- Helper constructors: `Text()`, `Paragraph()`, `Heading()`, `Bold()`, `Italic()`, `Image()`

**RTF reader (`reader/rtf/`)**
- Tree-based tokenizer: `{...}` groups are parsed as `Grouped` tokens containing their children, enabling correct group-scoped state
- Formatting: bold, italic, underline, strikethrough
- Paragraphs via `\par` and `\pard`
- Headings via `\outlinelevel` and Scrivener `<$Scr_H::N>` markers
- Lists with original markers preserved (`\pntext`, `\ls`, `\ilvl`) — supports bullets, numbered, roman numerals, and any custom marker
- Unicode via `\uN` and hex escapes via `\'XX`
- Extension groups `{\* ...}` correctly ignored
- Scrivener-specific metadata groups ignored (`fonttbl`, `colortbl`, `stylesheet`, `listtext`, etc.)

**Scrivener importer (`scriv/`)**
- Parses `.scrivx` XML (Scrivener 3 format with UUID-based file paths)
- Reads `Files/Data/<UUID>/content.rtf` for each document
- Preserves full binder tree with recursive children
- Respects `IncludeInCompile` flag
- Strips Scrivener style placeholders (`<$Scr_Cs::N>`, `<$Scr_Ps::N>`, `<$Scr_H::N>`) from imported content
- Skips `TrashFolder` items automatically

**Markdown writer (`writer/markdown/`)**
- Converts `ast.Document` to clean Markdown
- Preserves list markers as-is from the source document
- Supports all standard inline formatting and block elements

**HTML writer (`writer/html/`) — experimental**
- Basic conversion of `ast.Document` to HTML
- Not yet tested against real-world RTF output

### Known Limitations

- Images (`\pict`) are not yet extracted — skipped silently (planned for a future release)
- RTF tables render without cell separators — cells are concatenated
- Scrivener export (`.scriv` output) not yet implemented
- Markdown reader not yet implemented
- PDF, docx, and ePub writers not yet implemented

[0.1.0-alpha]: https://github.com/calmecac-dev/voluta/releases/tag/v0.1.0-alpha