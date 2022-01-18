package codegen

type Service struct {
	name      string
	contracts []Contract
}

func (s Service) Name() string {
	return s.name
}

func (s Service) NumContracts() int {
	return len(s.contracts)
}

func (s Service) Contract(i int) Contract {
	return s.contracts[i]
}
