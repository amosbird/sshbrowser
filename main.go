package main

import (
	"fmt"
	"net"
	"os"
)

func params() string {
	if len(os.Args) == 1 {
		return ""
	}
	return os.Args[1]
}

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9991")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}
	defer conn.Close()

	message := fmt.Sprintf("%s\n", params())

	buffer := []byte(message)

	_, err = conn.Write(buffer)
	if err != nil {
		fmt.Println("Error sending UDP packet:", err)
		return
	}
}
