package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// I need to know the followijng information
	// ip address of a node, and the port for listening
	ipAddress := "23.105.204.56"
	port := 8000

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ipAddress, port))
	if err != nil {
		fmt.Println("Something goes wrong", err)
		return
	}

	fmt.Fprintf(conn, "This is a test")

	status, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		fmt.Println("Something goes wrong", err)
	}

	fmt.Println(status)

	// Now try to talk to the node
	// Build a connection first
	// Send introduction message
	// Check the response
	//

	fmt.Println(ipAddress, port)

x}
