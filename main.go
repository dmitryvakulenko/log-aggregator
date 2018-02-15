package main

import (
	"net"
	"log-aggregator/database"
	//"log-aggregator/simhash"
	"log-aggregator/categorizer"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	log.Print("Starting listening logs...")
	udpAddr, err := net.ResolveUDPAddr("udp", ":3456")
	if err != nil {
		log.Fatalf("Error resolving address %v", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	defer conn.Close()

	if err != nil {
		log.Fatalf("Error listening port %v", err)
	}

	repo := database.Repository{}
	repo.Connect()
	defer repo.Disconnect()

	cat := categorizer.Categorizer{Storage: &repo}

	ch := make(chan []byte)
	go store(ch, cat)

	for {
		data, err := read(conn)
		if err == nil {
			ch <- data
		}
	}
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

func store(ch chan []byte, cat categorizer.Categorizer) {
	for {
		var r = &categorizer.LogRecord{}
		data := <-ch
		err := json.Unmarshal(data, r)
		if err != nil {
			log.Printf("Can't parse message %v", err)
			continue
		}

		cat.AddRecord(r)
	}
}