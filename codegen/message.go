package codegen

import (
	"fmt"
	"reflect"
	"strings"
)

type Field struct {
	name string
	typ  string
	tag  reflect.StructTag
}

func (f Field) Name() string {
	return f.name
}

func (f Field) Type() string {
	return f.typ
}

func (f Field) Tag(name string) string {
	return f.tag.Get(name)
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

func (m Message) PkgPath() string {
	return m.pkgPath
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
