package main

import (
	"net"
	"time"
	"log"
	"../conn"
	"os"
)

var (
	// dial the remote server, establish the control connection
	// TODO: separate the cltConn and the pxyConn
	ctlAddr = "127.0.0.1:9991"
	// dial the local server, establish the local connection
	locAddr = "127.0.0.1:22"
)

func main() {

	ctlConnCh := make(chan net.Conn)
	okCh := make(chan interface{})

	// connect to the server, and put the conn to the ctlConnCh
	ctlConnGetter := func() {
		ctlConn, err := net.Dial("tcp", ctlAddr)
		if err != nil {
			log.Printf("dial ctlAddr %v error: %v\n", ctlAddr, err)
			os.Exit(0)
		}
		log.Printf("dial ctlAddr %v \n", ctlAddr)
		ctlConnCh <- ctlConn
	}
	go ctlConnGetter()
	go func() {
		for {
			// get the ctlConn from ctlConnCh
			ctlConn := <-ctlConnCh

			// get the locConn
			locConn, err := net.Dial("tcp", locAddr)
			if err != nil {
				log.Printf("dial locAddr %v: %v error\n", locAddr, err)
				return
			}
			log.Printf("dial locAddr %v\n", locAddr)

			// transfer the data between server and the local network
			conn.Join(locConn, ctlConn)
			log.Printf("join ok!\n")
			okCh <- 1
		}
	}()

	// if join ok, reconnect the remote server
	go func() {
		for {
			<-okCh
			ctlConnGetter()
		}
	}()
	time.Sleep(10 * time.Minute)

}
