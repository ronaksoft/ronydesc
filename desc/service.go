package desc

type Service interface {
	Name() string
	Contracts() []Contract
}
