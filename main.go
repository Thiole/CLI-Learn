package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandList []Command

func main() {
	prefix := "⇒⇒"
	reader := bufio.NewReader(os.Stdin)
	/////////later will combine command list with a plugin command list for dynamic loading of commands and funcitonality

	/*	Connections.ConnectTCP{
		Address: "aardwolf.com",
		Port:    "23"}				*/
	commandList = []Command{
		{Name: "conn", Function: conn, Description: "Connect to new or existing connection"},
		{Name: "reverse", Function: reverse, Description: "Reverse and print the arguments"},
		{Name: "quit", Function: quit, Description: "Quit the program"},
		{Name: "help", Function: help, Description: "List All Commands"},
	}

	for {

		fmt.Print(prefix)
		//user input
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Split into parts, first is command, rest is args
		cmdParts := cmdSplit(input)
		command := cmdParts[0]
		args := cmdParts[1:]

		// loop through array of commands
		cmdFound := false //Check if a valid command is found in list
		for _, cmd := range commandList {
			if cmd.Name == command {
				cmd.Function(args)
				cmdFound = true
				break
			} else {

			}

		}
		if !cmdFound {
			fmt.Println("Invalid command")
		}
	}
}

type Command struct {
	Name        string
	Function    func([]string)
	Description string
}

func hello(args []string) {
	fmt.Println("Hello, world!")
	for i, arg := range args {
		fmt.Printf("arg[%d] = %s\n", i, arg)
	}
}

func conn(args []string) {
	Connect(args)
}

func reverse(args []string) {
	for i := len(args) - 1; i >= 0; i-- {
		fmt.Print(args[i] + " ")
	}
	fmt.Println()
}

func quit(args []string) {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func help(args []string) {
	fmt.Println("Commands:")
	for _, cmd := range commandList {
		fmt.Printf("  %-10s %s\n", cmd.Name, cmd.Description)
	}
}

func cmdSplit(cmd string) (cleanArgs []string) {

	args := strings.Split(cmd, " ")
	for _, arg := range args {
		if arg != "" {
			cleanArgs = append(cleanArgs, arg)
		}
	}
	return (cleanArgs)
}
