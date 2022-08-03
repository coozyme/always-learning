package after

import "fmt"

type IMessage interface {
	Create()
}

type Message struct {
	body string
}

func (m *Message) Create() string {
	m = &Message{
		body: "Hai Ary!",
	}

	res := m.body

	fmt.Println("Create Message")
	fmt.Printf("%+v\n", m)
	return res

}
