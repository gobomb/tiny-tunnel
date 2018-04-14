package main

import (
	"net"
	"time"
	"log"
	"../conn"
)

var (
	remoteAddress = "127.0.0.1:9991"
	localAddress  = "127.0.0.1:22"
)

func main() {
	// connect to the server
	remoteConn, err := net.Dial("tcp", remoteAddress)
	if err != nil {
		log.Printf("dial remoteAddress %v error: %v\n", remoteAddress, err)
	}
	log.Printf("dial remoteAddress %v \n", remoteAddress)

	// receive the notification from server
	//var sz int64
	//err = binary.Read(remoteConn, binary.LittleEndian, &sz)
	//if err != nil {
	//	log.Printf("read from remoteAddress %v error: %v\n", remoteAddress, err)
	//}
	//log.Printf("Reading message with length from remoteAddress %v : %d", remoteAddress, sz)
	//buffer := make([]byte, sz)
	//_, err = remoteConn.Read(buffer)
	//if err != nil {
	//	log.Printf("read from remoteAddress %v error: %v", remoteAddress, err)
	//}
	//log.Printf("Read message %v", buffer)

	// start transformation
	//if buffer[0] == 1 {
	// connect to  the local (behind-nat) server
	localConn, err := net.Dial("tcp", localAddress)
	if err != nil {
		log.Printf("dial localAddress %v: %v error\n", localAddress, err)
	}
	log.Printf("dial localAddress %v\n", localAddress)
	// transfer the data between server and the local network
	conn.Join(localConn, remoteConn)
	log.Printf("join ok!\n")
	time.Sleep(10 * time.Minute)
	//}
}
