package Connect

type StrRingBuffer struct {
	inputChannel  <-chan string
	outputChannel chan string
	logChannel    chan string
}

type IntRingBuffer struct {
	inputChannel  <-chan int
	outputChannel chan int
}

func NewStrRingBuffer(inputChannel <-chan []byte, outputChannel chan []byte, logChannel chan string) *StrRingBuffer {
	return &StrRingBuffer{inputChannel, outputChannel, logChannel}
}

func (r *StrRingBuffer) Run() {
	for v := range r.inputChannel {
		select {
		case r.outputChannel <- v:
		case r.logChannel <- v:
		default:
			<-r.outputChannel
			r.outputChannel <- v
			r.logChannel <- v
		}
	}
	close(r.outputChannel)
	close(r.logChannel)
}
