package codegen

type Message struct {
	name   string
	fields []Field
}

type Field struct {
	name string
	typ  string
}
