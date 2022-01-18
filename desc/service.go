package desc

import "github.com/ronaksoft/ronydesc/internal/gen"

type Service struct {
	Name      string
	Contracts []Contract
}

func (s Service) GetName() string {
	return s.Name
}

func (s Service) ForEachContract(f func(contract gen.Contract)) {
	for _, c := range s.Contracts {
		f(c)
	}
}

func (s Service) Register() Service {
	gen.Register(s)

	return s
}
