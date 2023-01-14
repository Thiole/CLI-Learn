package Connect

type Connection interface {
	Connect()
	Kill()
	Send()
	Receive()
	Log()
}

func Connect(args []string) {

}

func NewConnection(id string, name string) *TCPConnection {
	return &TCPConnection{
		ID:   id,
		Name: name,
		Generate: func() *StrRingBuffer {
			input := make(chan string)
			output := make(chan string, 30)
			log := make(chan string, 30)
			return NewStrRingBuffer(input, output, log)
		},
	}
}
