package codegen

type Contract struct {
	name   string
	input  Message
	output Message

	restMethod string
	restPath   string
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
	return c.restPath != ""
}

func (c Contract) Method() string {
	return c.restMethod
}

func (c Contract) Path() string {
	return c.restPath
}
