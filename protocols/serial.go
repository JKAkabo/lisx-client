package protocols

import (
	"go.bug.st/serial"
	"lisx/server"
	"log"
)

func InitSerial(
	portName string,
	baudRate int,
	parity serial.Parity,
	dataBits int,
	stopBits serial.StopBits,
	startDelim string,
	endDelim string,
) {
	mode := &serial.Mode{
		BaudRate: baudRate,
		Parity:   parity,
		DataBits: dataBits,
		StopBits: stopBits,
	}
	port, err := serial.Open(portName, mode)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer func(port serial.Port) {
		err := port.Close()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}(port)
	lBuf := LBuffer{StartDelim: startDelim, EndDelim: endDelim}
	for {
		lBuf.S += readSerial(&port)
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
}

func readSerial(port *serial.Port) string {
	var d string
	buf := make([]byte, 100)
	for {
		n, err := (*port).Read(buf)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		if n == 0 {
			break
		}
		d += string(buf[:n])
	}
	return d
}
