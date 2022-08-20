package valueobject

type Name struct {
	First, Last string
}

func (n *Name) Full() string {
	return n.First + " " + n.Last
}

func NewName(first, last string) *Name {
	return &Name{
		First: first,
		Last:  last,
	}
}
