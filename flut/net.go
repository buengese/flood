package flut

import (
	"log"
	"net"
)

// @speed: add some performance reporting mechanism on these functions when
//   called as goroutines

// bombAddress writes the given message via plain TCP to the given address,
// forever, as fast as possible.
func bombAddress(message []byte, address string) {
	i := 0
	for {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Println("bomb-error")
			i++
			if i >= 10 {
				log.Fatalf("dial failed after 10 retries, too many connections?")
			}
			continue
		}

		bombConn(message, conn)
		conn.Close()
	}
}

func bombConn(message []byte, conn net.Conn) {
	for {
		_, err := conn.Write(message)
		if err != nil {
			log.Println("conn-err")
			break
		}
	}
}
