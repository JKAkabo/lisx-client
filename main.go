package main

import (
	"github.com/joho/godotenv"
	"lisx/protocols"
	"lisx/server"
	"lisx/storage"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	accessToken, err := server.CreateAccessToken()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	storage.Token = accessToken.Token
	machine, err := server.GetMachine(accessToken.MachineID)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if machine.Protocol == "TCP" {
		if machine.Configuration.ServerMode {
			protocols.InitTcpServer(
				machine.Configuration.Port,
				machine.Configuration.StartDelimiter,
				machine.Configuration.EndDelimiter,
			)
		} else {
			protocols.InitTcpClient(
				machine.Configuration.IP,
				machine.Configuration.Port,
				machine.Configuration.StartDelimiter,
				machine.Configuration.EndDelimiter,
			)
		}
	} else if machine.Protocol == "SERIAL" {
		protocols.InitSerial(
			machine.Configuration.SerialPort,
			machine.Configuration.BaudRate,
			machine.Configuration.Parity,
			machine.Configuration.DataBits,
			machine.Configuration.StopBits,
			machine.Configuration.StartDelimiter,
			machine.Configuration.EndDelimiter,
		)
	}

}
