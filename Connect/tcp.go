package Connect

type TCPConnection struct {
	ID       string
	Name     string
	conn     *StrRingBuffer
	Generate func() *StrRingBuffer
}

func (c *TCPConnection) Open() {
	c.conn = c.Generate()
	go c.conn.Run()
}

func (c *TCPConnection) Close() {
	close(c.conn.outputChannel)
	close(c.conn.logChannel)
}

func (c *TCPConnection) Send(data string) {

}

func (c *TCPConnection) Receive() string {
	return <-c.conn.outputChannel
}

func (c *TCPConnection) Log() string {
	return <-c.conn.logChannel
}
