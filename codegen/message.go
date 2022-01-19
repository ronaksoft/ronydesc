package codegen

import (
	"fmt"
	"strings"
)

type Field struct {
	name string
	typ  string
}

func (f Field) Name() string {
	return f.name
}

func (f Field) Type() string {
	return f.typ
}

type Message struct {
	name    string
	pkgPath string
	fields  []Field
}

func (m Message) Name() string {
	return m.name
}

func (m Message) NumFields() int {
	return len(m.fields)
}

func (m Message) Field(i int) Field {
	return m.fields[i]
}

func (m Message) Fields() []Field {
	return m.fields
}

func (m Message) String() string {
	sb := strings.Builder{}
	sb.WriteRune('{')
	sb.WriteRune('\n')
	for _, f := range m.fields {
		sb.WriteString(fmt.Sprintf("\t%s: %s,\n", f.name, f.typ))
	}
	sb.WriteRune('}')

	return sb.String()
}
