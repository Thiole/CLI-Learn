package main

import (
	"Silint2/Connect"
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var sym string

type Command struct {
	Name        string
	Function    func([]string)
	Description string
}

func main() {
	prefix := sym + "⇒"
	reader := bufio.NewReader(os.Stdin)

	/// refactor into a command handler, so help, command list, etc is all in one place.
	// maybe multiple tiers of commands, better support for higher complexity runtimes later?
	commandList := []Command{
		{Name: "conn", Function: conn, Description: "Connect to new or existing connection"},
		{Name: "quit", Function: quit, Description: "Quit the program"},
		{Name: "snow", Function: snow, Description: "Service Now Management and Read"},
		{Name: "plugins", Function: plugins, Description: "Plugins Managment and Install"},
		{Name: "help", Function: help, Description: "List All Commands"},
	}
	///just so they are used temporarily
	for {

		fmt.Print(prefix, commandList)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		conn, err := net.Dial("aardmud.com", "23")

		inputChan := make(chan []byte)
		outputChan := make(chan []byte)
		// Start a goroutine to copy data from the connection to the input channel
		go func() {
			defer close(inputChan)
			for {
				buffer := make([]byte, 1024)
				n, err := conn.Read(buffer)
				if err != nil {
					if err != io.EOF {
						log.Println(err)
					}
					break
				}
				inputChan <- buffer[:n]
			}
		}()

		// Start a goroutine to copy data from the output channel to the connection
		go func() {
			defer close(outputChan)
			for data := range outputChan {
				_, err := conn.Write(data)
				if err != nil {
					log.Println(err)
					break
				}
			}
		}()

		logChan := make(chan string)
		Connect.NewStrRingBuffer(inputChan, outputChan, logChan)
		/*
			typestr := "tcp"
			typeargs := []string{"aardmud.com", "23"}
		*/

		//capture

	}
}
