package main

import (
	"net"
	"../conn"
	"time"
	"log"
)

var (
	// listen for the client, wait for the control connection
	// TODO: separate the cltConn and the pxyConn
	ctlAddr = "127.0.0.1:9991"
	// listen the request from the public network
	pubAddr = "127.0.0.1:9992"
)

func main() {
	ctlConnCh := make(chan net.Conn, 0)
	pubConnCh := make(chan net.Conn, 0)

	// listen the client connection
	ctlListener, err := net.Listen("tcp", ctlAddr)
	if err != nil {
		log.Printf("listen ctlAddr %v error: %v\n", ctlAddr, err)
		return
	}
	log.Printf("start listen ctlAddr %v\n", ctlAddr)
	// get the proxyConn
	go func() {
		for {
			ctlConn, err := ctlListener.Accept()
			if err != nil {
				log.Printf("accept ctlAddr  %v error: %v", ctlAddr, err)
			}
			log.Printf("accept ctlAddr %v\n", ctlAddr)
			ctlConnCh <- ctlConn
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
			ctlConn := <-ctlConnCh
			pubConn := <-pubConnCh

			// transfer the data between public network and the client
			conn.Join(ctlConn, pubConn)
			log.Printf("join ok!\n")
		}
	}()
	log.Printf("start!\n")
	time.Sleep(10 * time.Minute)
}
