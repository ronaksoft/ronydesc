package ex1

import "github.com/ronaksoft/ronydesc/desc"

var _ = desc.Service{
	Name: "serviceA",
	Contracts: []desc.Contract{
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
	},
}.Register()
