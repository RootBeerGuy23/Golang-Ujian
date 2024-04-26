package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8118")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		ClientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go HandleConn(ClientConn)
	}

}

func HandleConn(ClientConn net.Conn) {

	err := ClientConn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		panic(err)
	}
	var size uint32

	err = binary.Read(ClientConn, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	Bytmsg := make([]byte, size)

	_, err = ClientConn.Read(Bytmsg)
	if err != nil {
		panic(err)
	}
	RealMessage := string(Bytmsg)

	fmt.Println(RealMessage)

}
