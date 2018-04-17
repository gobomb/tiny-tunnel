package main

import (
	"net"
	"../conn"
	"time"
	"log"
)

//// TODO: add the cltConn
var (
	// listen for the client, wait for the proxy  connection
	pxyAddr = "127.0.0.1:9991"
	// listen the request from the public network
	pubAddr = "127.0.0.1:9992"
)

func main() {
	pxyConnCh := make(chan net.Conn, 0)
	pubConnCh := make(chan net.Conn, 0)

	// listen the client connection
	pxyListener, err := net.Listen("tcp", pxyAddr)
	if err != nil {
		log.Printf("listen pxyAddr %v error: %v\n", pxyAddr, err)
		return
	}
	log.Printf("start listen pxyAddr %v\n", pxyAddr)
	// get the proxyConn
	go func() {
		for {
			pxyConn, err := pxyListener.Accept()
			if err != nil {
				log.Printf("accept pxyAddr  %v error: %v", pxyAddr, err)
			}
			log.Printf("accept pxyAddr %v\n", pxyAddr)
			pxyConnCh <- pxyConn
		}
	}()

	// listen the public connection
	pubListener, err := net.Listen("tcp", pubAddr)
	if err != nil {
		log.Printf("listen pubAddr %v error: %v\n", pubAddr, err)
		return
	}
	log.Printf("start listen pubAddr %v\n", pubAddr)

	// get the pubConn
	go func() {
		for {
			pubConn, err := pubListener.Accept()
			if err != nil {
				log.Printf("accept pubAddr  %v error: %v\n", pubAddr, err)
			}
			log.Printf("accept pubAddr %v\n", pubAddr)
			pubConnCh <- pubConn
		}
	}()

	// join the pxyConn and pubConn
	go func() {
		for {
			pxyConn := <-pxyConnCh
			pubConn := <-pubConnCh

			// transfer the data between public network and the client
			conn.Join(pxyConn, pubConn)
			log.Printf("join ok!\n")
		}
	}()
	log.Printf("start!\n")
	time.Sleep(10 * time.Minute)
}
