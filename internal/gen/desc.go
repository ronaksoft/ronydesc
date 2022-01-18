package gen

type Service interface {
	GetName() string
	ForEachContract(f func(contract Contract))
}

type Contract interface {
	GetName() string
	GetInput() interface{}
	GetOutput() interface{}
	GetREST() REST
}

type REST interface {
	GetMethod() string
	GetPath() string
}
