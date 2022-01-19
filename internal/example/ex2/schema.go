package ex2

import (
	"github.com/ronaksoft/ronydesc/desc"
	"github.com/ronaksoft/ronydesc/internal/example/msg"
)

type ServiceB struct{}

func (s ServiceB) Name() string {
	return "serviceB"
}

func (s ServiceB) Contracts() []desc.Contract {
	return []desc.Contract{
		{
			Rests: []desc.REST{
				{
					Method: "GET",
					Path:   "/sum/:val1/:val2",
				},
				{
					Method: "POST",
					Path:   "/sum",
				},
			},
			Name:   "Sum",
			Input:  msg.SumRequest{},
			OutPut: msg.SumResponse{},
		},
	}
}
