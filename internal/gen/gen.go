package gen

var (
	registeredServices = map[string]Service{}
)

func Register(d Service) {
	registeredServices[d.GetName()] = d
}

func ForEachService(f func(service Service)) {
	for _, s := range registeredServices {
		f(s)
	}
}
