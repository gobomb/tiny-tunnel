package main

import (
	"net"
	"time"
	"log"
	"../conn"
	"os"
)

// TODO: add the cltConn
var (
	// dial the remote server, establish the proxy  connection
	pxyAddr = "127.0.0.1:9991"
	// dial the local server, establish the local connection
	locAddr = "127.0.0.1:22"
)

func main() {

	pxyConnCh := make(chan net.Conn)
	okCh := make(chan interface{})

	// connect to the server, and put the conn to the pxyConnCh
	pxyConnGetter := func() {
		pxyConn, err := net.Dial("tcp", pxyAddr)
		if err != nil {
			log.Printf("dial pxyAddr %v error: %v\n", pxyAddr, err)
			os.Exit(0)
		}
		log.Printf("dial pxyAddr %v \n", pxyAddr)
		pxyConnCh <- pxyConn
	}
	go pxyConnGetter()
	go func() {
		for {
			// get the pxyConn from pxyConnCh
			pxyConn := <-pxyConnCh

			// get the locConn
			locConn, err := net.Dial("tcp", locAddr)
			if err != nil {
				log.Printf("dial locAddr %v: %v error\n", locAddr, err)
				return
			}
			log.Printf("dial locAddr %v\n", locAddr)

			// transfer the data between server and the local network
			conn.Join(locConn, pxyConn)
			log.Printf("join ok!\n")
			okCh <- 1
		}
	}()

	// if join ok, reconnect the remote server
	go func() {
		for {
			<-okCh
			pxyConnGetter()
		}
	}()
	time.Sleep(10 * time.Minute)

}
