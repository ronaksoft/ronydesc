package codegen

type REST struct {
	Method string
	Path   string
}

type Contract struct {
	name   string
	input  Message
	output Message
	rests  []REST
}

func (c Contract) Name() string {
	return c.name
}

func (c Contract) Input() Message {
	return c.input
}

func (c Contract) Output() Message {
	return c.output
}

func (c Contract) IsREST() bool {
	return len(c.rests) > 0
}

func (c Contract) NumRests() int {
	return len(c.rests)
}

func (c Contract) Rest(i int) REST {
	return c.rests[i]
}
