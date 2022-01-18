package codegen

import (
	"reflect"

	"github.com/ronaksoft/ronydesc/internal/gen"
)

func Generate() *Arg {
	arg := &Arg{}
	gen.ForEachService(
		func(s gen.Service) {
			arg.extractService(s)
		},
	)

	return arg
}

type Arg struct {
	services  map[string]Service
	contracts map[string]Contract
	messages  map[string]Message
}

func (arg *Arg) extractService(s gen.Service) Service {
	svc := Service{
		name: s.GetName(),
	}
	s.ForEachContract(
		func(contract gen.Contract) {
			svc.contracts = append(svc.contracts, arg.extractContract(contract))
		},
	)

	arg.services[svc.name] = svc

	return svc
}

func (arg *Arg) extractContract(c gen.Contract) Contract {
	contract := Contract{
		name:   c.GetName(),
		input:  arg.extractMessage(c.GetInput()),
		output: arg.extractMessage(c.GetOutput()),
	}
	if c.GetREST() != nil {
		contract.restPath = c.GetREST().GetPath()
		contract.restMethod = c.GetREST().GetPath()
	}

	arg.contracts[contract.name] = contract
	return contract
}

func (arg *Arg) extractMessage(m interface{}) Message {
	rType := reflect.TypeOf(m)

	msg := Message{
		name: rType.Name(),
	}
	for i := 0; i < rType.NumField(); i++ {
		field := rType.Field(i)
		msg.fields = append(
			msg.fields,
			Field{
				name: field.Name,
				typ:  field.Type.String(),
			},
		)
	}

	arg.messages[msg.name] = msg

	return msg
}
