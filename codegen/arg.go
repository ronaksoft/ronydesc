package codegen

import (
	"reflect"

	"github.com/ronaksoft/ronydesc/desc"
)

type Arg struct {
	services  map[string]Service
	contracts map[string]Contract
	messages  map[string]Message

	PkgName string
}

func newArg(pkgName string) *Arg {
	return &Arg{
		PkgName:   pkgName,
		services:  map[string]Service{},
		contracts: map[string]Contract{},
		messages:  map[string]Message{},
	}
}

func (arg *Arg) extractService(s desc.IService) Service {
	svc := Service{
		pkgPath: reflect.TypeOf(s).PkgPath(),
		name:    s.Name(),
	}
	for _, c := range s.Contracts() {
		svc.contracts = append(svc.contracts, arg.extractContract(c))
	}

	arg.services[svc.name] = svc

	return svc
}

func (arg *Arg) extractContract(c desc.Contract) Contract {
	contract := Contract{
		name:   c.GetName(),
		input:  arg.extractMessage(c.GetInput()),
		output: arg.extractMessage(c.GetOutput()),
	}
	for _, r := range c.GetREST() {
		contract.rests = append(contract.rests, REST{
			Path:   r.Path,
			Method: r.Method,
		})
	}

	arg.contracts[contract.name] = contract

	return contract
}

func (arg *Arg) extractMessage(m interface{}) Message {
	rType := reflect.TypeOf(m)

	msg := Message{
		name:    rType.Name(),
		pkgPath: rType.PkgPath(),
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
