package codegen

import (
	"reflect"

	"github.com/ronaksoft/ronydesc/desc"
	"github.com/ronaksoft/ronydesc/internal/gen"
)

func Generate() *Arg {
	arg := newArg()
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

func newArg() *Arg {
	return &Arg{
		services:  map[string]Service{},
		contracts: map[string]Contract{},
		messages:  map[string]Message{},
	}
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
		if field.Type == reflect.TypeOf(desc.Message{}) {
			continue
		}
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

func (arg *Arg) NumServices() int {
	return len(arg.services)
}

func (arg *Arg) NumContracts() int {
	return len(arg.contracts)
}

func (arg *Arg) NumMessages() int {
	return len(arg.messages)
}

func (arg *Arg) Services() []Service {
	var arr = make([]Service, 0, len(arg.services))
	for _, v := range arg.services {
		arr = append(arr, v)
	}

	return arr
}

func (arg *Arg) Contracts() []Contract {
	var arr = make([]Contract, 0, len(arg.contracts))
	for _, v := range arg.contracts {
		arr = append(arr, v)
	}

	return arr
}

func (arg *Arg) Messages() []Message {
	var arr = make([]Message, 0, len(arg.messages))
	for _, v := range arg.messages {
		arr = append(arr, v)
	}

	return arr
}
