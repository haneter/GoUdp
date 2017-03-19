// UDPClient project main.go
package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var g_bufSize int

const op_start int = 1
const op_send int = 2
const op_finish int = 3

type dataPacket struct {
	opcode   int
	sequence int
	datasize int
	data     []byte
	md5      []byte
}

func main() {
	fmt.Println("Hello World!")

	var str_ipAddr string
	var int_ipPort int

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage : %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	tmp := strings.Contains(service, ":")
	if tmp == false {
		fmt.Fprintf(os.Stderr, "Usage : %s host:port", os.Args[0])
		os.Exit(1)
	}

	strtmp := strings.Split(service, ":")

	str_ipAddr = strtmp[0]
	int_ipPort, _ = strconv.Atoi(strtmp[1])

	fmt.Println("Result: ", str_ipAddr, " and ", int_ipPort)

	udpAddr, _ := net.ResolveUDPAddr("udp4", service)
	fmt.Println(udpAddr)
	//	checkError(err)

	conn, _ := net.DialUDP("udp", nil, udpAddr)
	//	checkError(err)

	//_, _ = conn.Write([]byte("anything"))
	_, _ = conn.Write([]byte("anything"))
	//	checkError(err)

	fmt.Println("1")
	//	var buf [512]byte
	//	n, _ := conn.Read(buf[0:])

	//	fmt.Println(string(buf[0:n]))

	client_openPort(int_ipPort)
	os.Exit(0)
}

func client_openPort(port int) {
	fmt.Println("client_openPort")

	port++
	service := ":" + strconv.Itoa(port)
	g_bufSize = 512 //Set default buf size

	udpAddr, _ := net.ResolveUDPAddr("udp4", service)
	//checkError(err)

	conn, _ := net.ListenUDP("udp", udpAddr)
	//checkError(err)

	go receiveLoop(conn)

}

func receiveLoop(conn *net.UDPConn) {
	for {
		handleclient(conn)
	}
}

func handleclient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(addr.IP)
	fmt.Println(buf)

}

func setDefaultBufSize(size int) {
	if size <= 0 {
		g_bufSize = 512
	} else {
		g_bufSize = size
	}
}

func makePacket(opcode int, data []byte) {

	switch opcode {
	case op_start:
	case op_send:
	case op_finish:
	}
}
