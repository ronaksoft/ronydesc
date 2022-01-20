package desc

type IService interface {
	Name() string
	Contracts() []Contract
}

type Service struct{}
