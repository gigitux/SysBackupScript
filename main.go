package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	endpoint := "insert_your_endpoint"
	token := "insert_your_token"
	isHTTPS := false
	message := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "backup",
		"params": map[string]string{
			"authToken": token,
		},
	}

	bytesRepresentation, err := json.Marshal(message)

	var tr *http.Transport

	if isHTTPS {
		tr = &http.Transport{}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s", endpoint), bytes.NewBuffer(bytesRepresentation))

	if err != nil {
		log.Println(err)
	}

	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
	}
}
