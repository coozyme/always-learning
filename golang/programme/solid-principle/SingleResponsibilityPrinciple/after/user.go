package after

type User struct {
	id    int
	name  string
	phone int
}

func (u *User) GetSender() *User {
	return &User{
		id:    1,
		name:  "Ary",
		phone: 6281,
	}
}

func (u *User) GetReceiver() *User {
	return &User{
		id:    2,
		name:  "Setya",
		phone: 6287,
	}
}
