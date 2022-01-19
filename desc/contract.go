package desc

type Contract struct {
	Name   string
	Input  IMessage
	OutPut IMessage
	Rests  []REST
}

func (c Contract) GetName() string {
	return c.Name
}

func (c Contract) GetInput() interface{} {
	return c.Input
}

func (c Contract) GetOutput() interface{} {
	return c.OutPut
}

func (c Contract) GetREST() []REST {
	return c.Rests
}

type REST struct {
	Method string
	Path   string
}
