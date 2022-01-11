package protocols

import (
	"bufio"
	"lisx/server"
	"log"
	"net"
	"strconv"
)

func InitTcpServer(port int, startDelim string, endDelim string) {
	l, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println("Waiting for connection on " + l.Addr().String())
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}(l)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Panicf("error: %v", err)
		}
		go func() {
			lBuf := LBuffer{StartDelim: startDelim, EndDelim: endDelim}
			r := bufio.NewReader(c)
			lBuf.S += readTcp(r)
			data := lBuf.GetData()
			if len(data) > 0 {
				go func() {
					err1 := server.CreateResults(data)
					if err1 != nil {
						log.Printf("error: %v", err)
					}
				}()
			}
		}()
	}
}

func InitTcpClient(ip string, port int, startDelim string, endDelim string) {
	c, err := net.Dial("tcp4", ip+":"+strconv.Itoa(port))
	if err != nil {
		log.Panicf("error: %v", err)
	}
	lBuf := LBuffer{StartDelim: startDelim, EndDelim: endDelim}
	r := bufio.NewReader(c)
	lBuf.S += readTcp(r)
	data := lBuf.GetData()
	if len(data) > 0 {
		go func() {
			err1 := server.CreateResults(data)
			if err1 != nil {
				log.Printf("error: %v", err)
			}
		}()
	}
}

func readTcp(r *bufio.Reader) string {
	var d string

	for {
		b, err := r.ReadByte()
		if err != nil {
			break
		}
		d += string(b)
	}
	return d
}
