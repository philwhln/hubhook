package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Payload struct {
	CallbackUrl string `json:"callback_url"`
	Repository  Repository
}

type Repository struct {
	RepoName string `json:"repo_name"`
	Status   string
}

type Callback struct {
	State       string
	Description string
	Content     string
	TargetUrl   string `json:"target_url"`
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LUTC)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var payload Payload
		var err error
		if payload, err = parsePayload(r); err != nil {
			logger.Println(err)
			return
		}
		fmt.Fprintf(w, "%s\n", payload)
		if err := sendCallback(payload.CallbackUrl, err); err != nil {
			logger.Printf("Failed to send callback to %s : %s\n", payload.CallbackUrl, err)
		}
		logger.Printf("repo:%s status:%s\n", payload.Repository.RepoName, payload.Repository.Status)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parsePayload(r *http.Request) (p Payload, err error) {
	var b []byte
	if b, err = ioutil.ReadAll(r.Body); err != nil {
		fmt.Println(err)
		return p, err
	}
	if err = json.Unmarshal(b, &p); err != nil {
		return p, err
	}
	return p, nil
}

func sendCallback(url string, processingErr error) error {
	var callback Callback
	if processingErr == nil {
		callback.State = "success"
	} else {
		callback.State = "error"
	}
	body, err := json.Marshal(callback)
	if err != nil {
		return err
	}
	http.Post(url, "application/json", bytes.NewReader(body))
	return nil
}
