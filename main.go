/* Copyright (c) 2022, Oracle and/or its affiliates.
Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type config struct {
	address string
	port    int
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		buffer := make([]byte, 2048)
		s, err := reader.Read(buffer)
		if err != nil {
			log.Fatalln("Unable to read data", err)
		}
		log.Println(string(buffer[0:s]))
	}

}

func main() {
	var cfg config
	// read the values for address and port from command line parameters
	flag.StringVar(&cfg.address, "address", "", "Listening address")
	flag.IntVar(&cfg.port, "port", 20080, "Listening port")

	//form the address
	address := fmt.Sprintf("%s:%d", cfg.address, cfg.port)

	// start a listener on config parameters
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Unable to bind address,port", err)
	}

	// waiting for connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		log.Printf("Connection established on %q\n", address)

		go handleConn(conn)
	}

}
