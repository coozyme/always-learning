package valueobject

type Password struct {
	Salt, Hash string
}

func (p *Password) SetHash(password string) {
	p.Hash = password
}

func (p *Password) Verify(password string) bool {
	if p.Hash != password {
		return false
	}
	return p.Hash == password
}

func NewPassword(salt, password string) *Password {
	return &Password{
		Salt: salt,
		Hash: password,
	}
}
