// The MIT License (MIT)
//
// Copyright (c) 2016-2020 Containous SAS; 2020-2022 Traefik Labs
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/importer"
	"go/token"
	"go/types"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"golang.org/x/tools/imports"
)

const closingLines = `

// vim: set nowrap :
`

// File a kind of AST element that represents a file.
type File struct {
	Package  string
	Imports  []string
	Elements []Element
}

// Element is a simplified version of a symbol.
type Element struct {
	Name  string
	Value string
}

// Centrifuge a centrifuge.
// Generate Go Structures from Go structures.
type Centrifuge struct {
	IncludedImports []string
	ExcludedTypes   []string
	ExcludedFiles   []string

	TypeCleaner    func(types.Type, string) string
	PackageCleaner func(string) string

	rootPkg string
	fileSet *token.FileSet
	pkg     *types.Package
}

// NewCentrifuge creates a new Centrifuge.
func NewCentrifuge(rootPkg string) (*Centrifuge, error) {
	fileSet := token.NewFileSet()

	pkg, err := importer.ForCompiler(fileSet, "source", nil).Import(rootPkg)
	if err != nil {
		return nil, err
	}

	return &Centrifuge{
		fileSet: fileSet,
		pkg:     pkg,
		rootPkg: rootPkg,

		TypeCleaner: func(typ types.Type, _ string) string {
			return typ.String()
		},
		PackageCleaner: func(s string) string {
			return s
		},
	}, nil
}

// Run runs the code extraction and the code generation.
func (c Centrifuge) Run(dest string, pkgName string) error {
	files, err := c.run(c.pkg.Scope(), c.rootPkg, pkgName)
	if err != nil {
		return err
	}

	err = fileWriter{baseDir: dest}.Write(files)
	if err != nil {
		return err
	}

	for _, p := range c.pkg.Imports() {
		if contains(c.IncludedImports, p.Path()) {
			fls, err := c.run(p.Scope(), p.Path(), p.Name())
			if err != nil {
				return err
			}

			err = fileWriter{baseDir: filepath.Join(dest, p.Name())}.Write(fls)
			if err != nil {
				return err
			}
		}
	}

	return err
}

func (c Centrifuge) run(sc *types.Scope, rootPkg string, pkgName string) (map[string]*File, error) {
	files := map[string]*File{}

	for _, name := range sc.Names() {
		if contains(c.ExcludedTypes, name) {
			continue
		}

		o := sc.Lookup(name)
		if !o.Exported() {
			continue
		}

		filename := filepath.Base(c.fileSet.File(o.Pos()).Name())
		if contains(c.ExcludedFiles, path.Join(rootPkg, filename)) {
			continue
		}

		// I can't imagine we're going to need any mocked, testing structures
		if strings.HasPrefix(filename, "mock_") {
			log.Printf("skipping due to excluded prefix: %s", filename)
			continue
		}

		fl, ok := files[filename]
		if !ok {
			files[filename] = &File{Package: pkgName}
			fl = files[filename]
		}

		elt := Element{
			Name: name,
		}

		switch ob := o.(type) {
		case *types.TypeName:

			switch obj := ob.Type().(*types.Named).Underlying().(type) {
			case *types.Struct:
				elt.Value = c.writeStruct(name, obj, rootPkg, fl)

			case *types.Map:
				elt.Value = fmt.Sprintf("type %s map[%s]%s\n", name, obj.Key().String(), c.TypeCleaner(obj.Elem(), rootPkg))

			case *types.Slice:
				elt.Value = fmt.Sprintf("type %s []%v\n", name, c.TypeCleaner(obj.Elem(), rootPkg))

			case *types.Basic:
				elt.Value = fmt.Sprintf("type %s %v\n", name, obj.Name())

			default:
				log.Printf("OTHER TYPE::: %s %T\n", name, o.Type().(*types.Named).Underlying())
				continue
			}

		default:
			log.Printf("OTHER::: %s %T\n", name, o)
			continue
		}

		if len(elt.Value) > 0 {
			fl.Elements = append(fl.Elements, elt)
		}
	}

	return files, nil
}

func (c Centrifuge) writeStruct(name string, obj *types.Struct, rootPkg string, elt *File) string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("type %s struct {\n", name))

	for i := 0; i < obj.NumFields(); i++ {
		field := obj.Field(i)

		if !field.Exported() {
			continue
		}

		fPkg := c.PackageCleaner(extractPackage(field.Type()))
		if fPkg != "" && fPkg != rootPkg {
			elt.Imports = append(elt.Imports, fPkg)
		}

		fType := c.TypeCleaner(field.Type(), rootPkg)

		if field.Embedded() {
			b.WriteString(fmt.Sprintf("\t%s\n", fType))
			continue
		}

		// skip if we have a "-" toml tag
		values, _ := lookupTagValue(obj.Tag(i), "toml")
		if len(values) > 0 && values[0] == "-" {
			continue
		}

		b.WriteString(fmt.Sprintf("\t%s %s", field.Name(), fType))

		tags := obj.Tag(i)
		if tags != "" {
			b.WriteString(" `" + tags + "`")
		}

		b.WriteString("\n")
	}

	b.WriteString("}\n")

	return b.String()
}

func lookupTagValue(raw, key string) ([]string, bool) {
	value, ok := reflect.StructTag(raw).Lookup(key)
	if !ok {
		return nil, ok
	}

	values := strings.Split(value, ",")

	if len(values) < 1 {
		return nil, true
	}

	return values, true
}

func extractPackage(t types.Type) string {
	switch tu := t.(type) {
	case *types.Named:
		return tu.Obj().Pkg().Path()

	case *types.Slice:
		if v, ok := tu.Elem().(*types.Named); ok {
			return v.Obj().Pkg().Path()
		}
		return ""

	case *types.Map:
		if v, ok := tu.Elem().(*types.Named); ok {
			return v.Obj().Pkg().Path()
		}
		return ""

	case *types.Pointer:
		return extractPackage(tu.Elem())

	default:
		return ""
	}
}

func contains(values []string, value string) bool {
	for _, val := range values {
		if val == value {
			return true
		}
	}

	return false
}

type fileWriter struct {
	baseDir string
}

func (f fileWriter) Write(files map[string]*File) error {
	err := os.MkdirAll(f.baseDir, 0o755)
	if err != nil {
		return err
	}

	for name, file := range files {
		err = f.writeFile(name, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f fileWriter) writeFile(name string, desc *File) error {
	if len(desc.Elements) == 0 {
		return nil
	}

	filename := filepath.Join(f.baseDir, name)

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer func() { _ = file.Close() }()

	b := bytes.NewBufferString("package ")
	b.WriteString(desc.Package)
	b.WriteString("\n")
	b.WriteString("// Code generated by centrifuge. DO NOT EDIT.\n")

	b.WriteString("\n")
	f.writeImports(b, desc.Imports)
	b.WriteString("\n")

	for _, elt := range desc.Elements {
		b.WriteString(elt.Value)
		b.WriteString("\n")
	}

	b.WriteString(closingLines)

	// gofmt
	source, err := format.Source(b.Bytes())
	if err != nil {
		log.Println(b.String())
		return fmt.Errorf("failed to format sources: %w", err)
	}

	// goimports
	process, err := imports.Process(filename, source, nil)
	if err != nil {
		log.Println(string(source))
		return fmt.Errorf("failed to format imports: %w", err)
	}

	_, err = file.Write(process)
	if err != nil {
		return err
	}

	return nil
}

func (f fileWriter) writeImports(b io.StringWriter, imports []string) {
	if len(imports) == 0 {
		return
	}

	uniq := map[string]struct{}{}

	sort.Strings(imports)

	_, _ = b.WriteString("import (\n")
	for _, s := range imports {
		if _, exist := uniq[s]; exist {
			continue
		}

		uniq[s] = struct{}{}

		_, _ = b.WriteString(fmt.Sprintf(`	"%s"`+"\n", s))
	}

	_, _ = b.WriteString(")\n")
}
