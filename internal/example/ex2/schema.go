package ex2

import "github.com/ronaksoft/ronydesc/desc"

var _ = desc.Service{
	Name: "serviceB",
	Contracts: []desc.Contract{
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
	},
}.Register()
