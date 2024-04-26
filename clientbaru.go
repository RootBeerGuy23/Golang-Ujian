package main

import (
	"encoding/binary"
	"net"
	"time"
)

func main() {
	SendMessage("halo")
}

func SendMessage(msg string) {
	ServerConn, err := net.DialTimeout("tcp", "127.0.0.1:8118", 3*time.Second)
	if err != nil {
		panic(err)
	}
	defer ServerConn.Close()

	err = binary.Write(ServerConn, binary.LittleEndian, uint32(len(msg)))
	if err != nil {
		panic(err)
	}

	_, err = ServerConn.Write([]byte(msg))
	if err != nil {
		panic(err)
	}

}
