package before

import "fmt"

type IMessage interface {
	Create()
	Send()
}

type Message struct {
	body     string
	sender   int
	receiver int
}

func (m *Message) Create() *Message {
	msg := Message{
		body:     "Hi, aryy",
		sender:   81231,
		receiver: 329023,
	}

	fmt.Println("Create Message")
	fmt.Printf("%+v\n", msg)

	return &msg
}

func (m *Message) Send(msg *Message) *Message {
	fmt.Println("Send Message")
	fmt.Println("Sender:", m.sender)
	fmt.Println("Receiver:", m.receiver)
	return msg
}
