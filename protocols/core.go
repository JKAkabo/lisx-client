package protocols

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type LBuffer struct {
	S string
	StartDelim string
	EndDelim string
}

func (b *LBuffer) GetData() (data []string) {
	for {
		i := strings.Index(b.S, b.StartDelim)
		j := strings.Index(b.S, b.EndDelim)
		if i == -1 || j == -1 {
			return data
		}
		j = j + len(b.EndDelim)
		data = append(data, (b.S)[i:j])
		b.S = (b.S)[j:]
	}
}

func Persist(data []string) {
	j, err := json.Marshal(data)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	req, err := http.NewRequest(http.MethodPut, "http://api.lisx.akabo.io/results", bytes.NewBuffer(j))
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	log.Printf("response status: %s", res.Status)
}