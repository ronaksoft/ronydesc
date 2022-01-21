package ex2

import (
	"github.com/ronaksoft/ronydesc/desc"
)

type ServiceB struct {
	desc.Service
}

func (s ServiceB) Name() string {
	return "ServiceB"
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
			Input:  SumRequest{},
			OutPut: SumResponse{},
		},
	}
}

type SumRequest struct {
	desc.Message

	Val1 int32 `json:"val1"`
	Val2 int32 `json:"val2"`
}

type SumResponse struct {
	desc.Message

	Sum int32 `json:"sum"`
}
