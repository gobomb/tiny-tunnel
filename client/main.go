package main

import (
	"net"
	"time"
	"log"
	"../conn"
	"os"
)

var (
	publicAddress = "127.0.0.1:9991"
	localAddress  = "127.0.0.1:22"
)

func main() {

	pubConnCh := make(chan net.Conn)
	okCh := make(chan interface{})

	// connect to the server, and put the conn to the pubConnCh
	pubConnGeter := func() {
		publicConn, err := net.Dial("tcp", publicAddress)
		if err != nil {
			log.Printf("dial publicAddress %v error: %v\n", publicAddress, err)
			os.Exit(0)
		}
		log.Printf("dial publicAddress %v \n", publicAddress)
		pubConnCh <- publicConn
	}
	go pubConnGeter()
	go func() {
		for {
			// get the publicConn from pubConnCh
			publicConn := <-pubConnCh

			// get the localConn
			localConn, err := net.Dial("tcp", localAddress)
			if err != nil {
				log.Printf("dial localAddress %v: %v error\n", localAddress, err)
				return
			}
			log.Printf("dial localAddress %v\n", localAddress)

			// transfer the data between server and the local network
			conn.Join(localConn, publicConn)
			log.Printf("join ok!\n")
			okCh <- 1
		}
	}()

	// if join ok, reconnect the public server
	go func() {
		for {
			<-okCh
			pubConnGeter()
		}
	}()
	time.Sleep(10 * time.Minute)

}
