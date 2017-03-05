// UDPClient project main.go
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage : %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	udpAddr, _ := net.ResolveUDPAddr("udp4", service)
	fmt.Println(udpAddr)
	//	checkError(err)

	conn, _ := net.DialUDP("udp", nil, udpAddr)
	//	checkError(err)

	//_, _ = conn.Write([]byte("anything"))
	_, _ = conn.Write([]byte("anything"))
	//	checkError(err)

	var buf [512]byte
	n, _ := conn.Read(buf[0:])

	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
