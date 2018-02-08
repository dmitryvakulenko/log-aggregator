package main

import (
	"net"
	"fmt"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":3456")
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}

	for {
		display(conn)
	}

	conn.Close()
}


func display(conn *net.UDPConn) {

	var buf [2048]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Println("Error Reading")
		return
	} else {
		fmt.Println(string(buf[0:n]))
		fmt.Println("Package Done")
	}

}