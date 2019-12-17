package internal

import (
	"cloud.google.com/go/spanner/spansql"
)

type Table struct {
	spansql.CreateTable

	indexes  []spansql.CreateIndex
	children []Table
}

type TemplateData struct {
	Table  *Table
	Tables []Table
}
