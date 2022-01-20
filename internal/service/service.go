package service

import (
	"github.com/ronaksoft/ronykit"
	"github.com/ronaksoft/ronykit/std/bundle/rest"
	"github.com/ronaksoft/ronykit/std/contract"
	"github.com/ronaksoft/ronykit/std/service"
)

/*
   This Code is Auto-Generated by RonyDESC. Please DO NOT EDIT.
   Generated at:
*/

// IServiceA define the interface which MUST be implemented.
type IServiceA interface {
	Echo(in *EchoRequest, out *EchoResponse)
}

// serviceaWrapper
type serviceaWrapper struct {
	svc IServiceA
}

func (x serviceaWrapper) echoContract() ronykit.Contract {
	return contract.New().
		SetSelector(
			rest.Route("GET", "/echo/:randomID").
				WithFactory(
					func() ronykit.Message {
						return &EchoRequest{}
					},
				),
		).
		SetHandler(func(ctx *ronykit.Context) {
			req := ctx.In().GetMsg().(*EchoRequest)
			res := &EchoResponse{}

			x.svc.Echo(req, res)

			out := ctx.Out()
			ctx.In().WalkHdr(
				func(key string, val string) bool {
					out.SetHdr(key, val)

					return true
				},
			)

			out.
				SetMsg(res).
				Send()
		})
}

func (x serviceaWrapper) Service() ronykit.Service {
	return service.New("ServiceA").
		AddContract(x.echoContract()).
		SetPreHandlers()
}
