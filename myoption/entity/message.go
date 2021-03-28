package entity

//  选项模式
type Message struct {
	id      int
	name    string
	phone   int
	address string
}

type Option func(msg *Message)

var DEFAULT_MESSAGE = Message{id: -1, name: "-1", address: "-1", phone: -1}

func WithID(id int) Option {
	return func(m *Message) {
		m.id = id
	}
}
func WithName(name string) Option {
	return func(m *Message) {
		m.name = name
	}
}
func WithAddress(addr string) Option {
	return func(m *Message) {
		m.address = addr
	}
}
func WithPhone(phone int) Option {
	return func(m *Message) {
		m.phone = phone
	}
}

func NewByOption(opts ...Option) Message {
	msg := DEFAULT_MESSAGE
	for _, f := range opts {
		f(&msg)
	}
	return msg
}

func NewRequireIDByOption(id int, opts ...Option) Message {
	msg := DEFAULT_MESSAGE
	msg.id = id
	for _, o := range opts {
		o(&msg)
	}
	return msg
}
