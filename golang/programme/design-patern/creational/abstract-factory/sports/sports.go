package sports

type SportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

type IShoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type IShirt interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}
