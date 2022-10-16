package arms

type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}
