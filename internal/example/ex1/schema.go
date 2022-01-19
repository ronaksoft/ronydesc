package ex1

import (
	"github.com/ronaksoft/ronydesc/desc"
	"github.com/ronaksoft/ronydesc/internal/example/msg"
)

type ServiceA struct{}

func (s ServiceA) Name() string {
	return "serverA"
}

func (s ServiceA) Contracts() []desc.Contract {
	return []desc.Contract{
		{
			Rests: []desc.REST{
				{
					Method: "GET",
					Path:   "/echo/:randomID",
				},
			},
			Name:   "Echo",
			Input:  msg.EchoRequest{},
			OutPut: msg.EchoResponse{},
		},
	}
}
