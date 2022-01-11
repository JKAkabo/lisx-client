package server

import (
	"bytes"
	"encoding/json"
	"go.bug.st/serial"
	"lisx/storage"
	"net/http"
)

type AccessToken struct {
	Token string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
	UserID string `json:"userID"`
	MachineID string `json:"machineID"`
}

type Machine struct {
	Name string `json:"name"`
	Protocol string `json:"protocol"`
	Configuration Configuration `json:"configuration"`
}

type Configuration struct {
	IP string `json:"ip"`
	Port int `json:"port"`
	ServerMode bool `json:"serverMode"`
	SerialPort string `json:"serialPort"`
	BaudRate int `json:"baudRate"`
	Parity serial.Parity `json:"parity"`
	DataBits int `json:"DataBits"`
	StopBits serial.StopBits `json:"stopBits"`
	StartDelimiter string `json:"startDelimiter"`
	EndDelimiter string `json:"endDelimiter"`
}

func CreateAccessToken() (accessToken AccessToken, err error) {
	body, err := json.Marshal(map[string]string{"username": storage.APIUsername, "password": storage.APIPassword})
	if err != nil {
		return accessToken, err
	}
	client := http.Client{}
	res, err := client.Post(storage.APIServer + "/access-tokens", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return accessToken, err
	}
	err = json.NewDecoder(res.Body).Decode(&accessToken)
	if err != nil {
		return accessToken, err
	}
	return accessToken, nil
}

func GetMachine(machineID string) (machine Machine, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", storage.APIServer + "/machines/" + machineID, nil)
	if err != nil {
		return machine, err
	}
	req.Header.Set("Authorization", "Bearer " + storage.Token)
	res, err := client.Do(req)
	if err != nil {
		return machine, err
	}
	err = json.NewDecoder(res.Body).Decode(&machine)
	if err != nil {
		return machine, err
	}
	return machine, nil
}

func CreateResults(data []string) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	client := http.Client{}
	req, err := http.NewRequest("GET", storage.APIServer + "/results", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer " + storage.Token)
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
