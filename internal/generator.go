package internal

import (
	"bytes"
	"text/template"
)

type GeneratorOption struct {
	Out          string
	Suffix       string
	Single       bool
	Template     string
	IgnoreFields []string
	IgnoreTables []string
}

type Generator struct {
	source       Source
	out          string
	single       bool
	template     string
	ignoreFields []string
	ignoreTables []string
}

func NewGenerator(source Source, option GeneratorOption) *Generator {
	return nil
}

// func (g *Generator) Generate() error {
// 	return nil
// }

// func (g *Generator) generate() {
// 	var buf bytes.Buffer

// 	g.template().Exec

// }

// func (g *Generator) template() error {

// }

// type Template struct {
// 	path string
// }

// func NewTemplate(path string) {

// }
