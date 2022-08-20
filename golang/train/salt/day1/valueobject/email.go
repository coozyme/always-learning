package valueobject

type Email struct {
	Local, Domain string
}

func (e *Email) ToString() string {
	return e.Local + "@" + e.Domain
}

func NewEmail(local, domain string) *Email {
	return &Email{
		Local:  local,
		Domain: domain,
	}
}
