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
	ctlAddr = "127.0.0.1:9000"
	// dial the remote server, establish the proxy  connection
	pxyAddr = "127.0.0.1:9991"
	// dial the local server, establish the local connection
	locAddr = "127.0.0.1:22"

	pxyConnCh chan net.Conn
	okCh      chan interface{}
	ctlConnCh chan net.Conn
)

func ctlConnGetter() {
	ctlConnCh = make(chan net.Conn, 0)

	ctlConn, err := net.Dial("tcp", ctlAddr)
	if err != nil {
		log.Printf("dial ctlAddr %v error: %v\n", ctlAddr, err)
		os.Exit(0)
	}
	log.Printf("dial ctlAddr %v \n", ctlAddr)
	ctlConnCh <- ctlConn

	// get the msg from the ctlConn
	go func() {
		// the msg "establishPxyConn
		go pxyConnGetter()
	}()
}

// connect to the server, and put the conn to the pxyConnCh
func pxyConnGetter() {
	pxyConn, err := net.Dial("tcp", pxyAddr)
	if err != nil {
		log.Printf("dial pxyAddr %v error: %v\n", pxyAddr, err)
		os.Exit(0)
	}
	log.Printf("dial pxyAddr %v \n", pxyAddr)
	pxyConnCh <- pxyConn
}

// if get the pxyConn, dial to the local address
func locAddrGetter() {
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
}


func main() {

	pxyConnCh = make(chan net.Conn)
	okCh = make(chan interface{})

	go pxyConnGetter()
	go locAddrGetter()

	// if join ok, reconnect the remote server
	go func() {
		for {
			<-okCh
			pxyConnGetter()
		}
	}()
	time.Sleep(10 * time.Minute)

}
