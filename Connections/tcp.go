package Connections

import (
	"net"
)

type ConnectTCP struct {
	Name    string
	Address string
	Port    string
}

func (r *ConnectTCP) Connect() error {

	addr, err := net.ResolveTCPAddr("tcp", r.Address+":"+r.Port)
	if err != nil {
		return err
	}

	// Establish the connection
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send and receive data over the connection
	// ...

	return nil
}

func (r *ConnectTCP) Close() error {
	return nil
}

func (r *ConnectTCP) Receive() error {
	return nil
}

func (r *ConnectTCP) Send() error {
	return nil
}
