package main

import (
	"net"
	"fmt"
	"encoding/json"
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

	ch := make(chan []byte)
	go store(ch)

	for {
		data, err := read(conn)
		if err == nil {
			ch <- data
		}
	}

	conn.Close()
}

func read(conn *net.UDPConn) ([]byte, error) {
	buf := make([]byte, 65535)
	length, err := conn.Read(buf)

	if err != nil {
		return nil, err
	} else {
		return buf[0:length], nil
	}
}

func store(ch chan []byte) {
	for {
		var r = &LogRecord{}
		data := <-ch
		err := json.Unmarshal(data, r)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(r.SimHash())
	}
}