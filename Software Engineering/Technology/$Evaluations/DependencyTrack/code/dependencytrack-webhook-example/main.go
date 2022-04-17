package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("DT_WEBHOOK_PORT")
	if port == "" {
		port = "3001"
	}
	http.HandleFunc("/webhook", webhookHandler)
	log.Printf("running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type NotificationPayload struct {
	Notification struct {
		Level     string      `json:"level"`
		Scope     string      `json:"scope"`
		Group     string      `json:"group"`
		Timestamp string      `json:"timestamp"`
		Title     string      `json:"title"`
		Content   string      `json:"content"`
		Subject   interface{} `json:"subject"`
	} `json:"notification"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var n NotificationPayload
	err = json.Unmarshal(body, &n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Printf("Notification: %+v\n", n)
	w.WriteHeader(http.StatusOK)
}
