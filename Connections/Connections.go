package Connections

type Connection interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}
