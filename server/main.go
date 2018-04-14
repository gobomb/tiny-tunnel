package main

import (
	"net"
	"../conn"
	"time"
	"log"
)

var (
	proxyAddress  = "127.0.0.1:9991"
	publicAddress = "127.0.0.1:9992"
)

func main() {
	connCh := make(chan net.Conn, 0)
	connCh2 := make(chan net.Conn, 0)

	// listen the client connection
	proxyListener, err := net.Listen("tcp", proxyAddress)
	if err != nil {
		log.Printf("listen proxyAddress %v error: %v\n", proxyAddress, err)
	}
	log.Printf("start listen proxyAddress %v\n", proxyAddress)
	go func() {
		for {
			proxyConn, err := proxyListener.Accept()
			if err != nil {
				log.Printf("accept proxyAddress  %v error: %v", proxyAddress, err)
			}
			log.Printf("accept proxyAddress %v\n", proxyAddress)
			//go func() {
			connCh <- proxyConn
			//}()
		}
	}()

	// listen the public connection
	publicListener, err := net.Listen("tcp", "127.0.0.1:9992")
	if err != nil {
		log.Printf("listen publicAddress %v error: %v\n", publicAddress, err)
	}
	log.Printf("start listen publicAddress %v\n", publicAddress)
	go func() {
		for {
			publicConn, err := publicListener.Accept()
			if err != nil {
				log.Printf("accept publicAddress  %v error: %v\n", publicAddress, err)
			}
			log.Printf("accept publicAddress %v\n", publicAddress)
			connCh2<-publicConn
		}
	}()
	go func() {
		for {
			proxyConn := <-connCh
			publicConn := <-connCh2
			//msg := []byte{1}
			//err = binary.Write(proxyConn, binary.LittleEndian, int64(len(msg)))
			//if err != nil {
			//	log.Printf("write to %v error: %v\n", proxyAddress, err)
			//}
			//
			//// send the msg to notify the client: start transformation
			//n, err := proxyConn.Write(msg)
			//if err != nil {
			//	log.Printf("write to %v error: %v\n", proxyAddress, err)
			//}
			//log.Printf("write %v byte\n", n)

			// transfer the data between public network and the client
			conn.Join(proxyConn, publicConn)
			log.Printf("join ok!\n")
		}
	}()
	log.Printf("start!\n")
	time.Sleep(10 * time.Minute)
}
