package codegen

type Service struct {
	name      string
	contracts []Contract
	pkgPath   string
}

func (s Service) Name() string {
	return s.name
}

func (s Service) PkgPath() string {
	return s.pkgPath
}

func (s Service) NumContracts() int {
	return len(s.contracts)
}

func (s Service) Contract(i int) Contract {
	return s.contracts[i]
}

func (s Service) Contracts() []Contract {
	return s.contracts
}
