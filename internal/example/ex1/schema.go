package ex1

import (
	"github.com/ronaksoft/ronydesc/desc"
)

type ServiceA struct {
	desc.Service
}

func (s ServiceA) Name() string {
	return "serviceA"
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
			Input:  EchoRequest{},
			OutPut: EchoResponse{},
		},
	}
}

type EchoRequest struct {
	desc.Message

	RandomID  string `json:"randomId"`
	RandomInt int32  `json:"randomInt"`
}

type EchoResponse struct {
	desc.Message

	RandomID  string `json:"randomId"`
	RandomInt int32  `json:"randomInt"`
}
