package desc

import "github.com/ronaksoft/ronydesc/internal/gen"

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

func (c Contract) GetREST() []gen.REST {
	var arr = make([]gen.REST, 0, len(c.Rests))
	for _, v := range c.Rests {
		arr = append(arr, v)
	}

	return arr
}

type REST struct {
	Method string
	Path   string
}

func (r REST) GetMethod() string {
	return r.Method
}

func (r REST) GetPath() string {
	return r.Path
}
