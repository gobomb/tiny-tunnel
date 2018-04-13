package conn

import (
	"sync"
	"net"
	"log"
	"io"
)
func Join(rawConn1, rawConn2 net.Conn) {
	log.Printf("start join!\n")
	var wait sync.WaitGroup

	pipe := func(to net.Conn, from net.Conn, bytesCopied *int64) {
		defer to.Close()
		defer from.Close()
		defer wait.Done()

		var err error
		*bytesCopied, err = io.Copy(to, from)
		if err != nil {
			log.Printf("io error: %v\n", err)
		}
	}

	wait.Add(2)
	var fromBytes, toBytes int64
	go pipe(rawConn1, rawConn2, &fromBytes)
	go pipe(rawConn2, rawConn1, &toBytes)
	wait.Wait()
}
