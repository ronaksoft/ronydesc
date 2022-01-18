package desc

import "github.com/ronaksoft/ronydesc/internal/gen"

type Contract struct {
	Name   string
	Input  IMessage
	OutPut IMessage
	Rest   *REST
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

func (c Contract) GetREST() gen.REST {
	return c.Rest
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
