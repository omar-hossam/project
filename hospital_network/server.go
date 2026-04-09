package main

import (
	"fmt"
	"net/http"
)

type EmergencyMessage struct {
	RawText      string
	Source       string
	HouseNumber  string
	Priority     string
	Injury       string
}

func main() {
	http.HandleFunc("/emergency", emergencyHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Hospital server listening on :8080")
	fmt.Println("Try: curl -X POST http://localhost:8080/emergency -d 'Help from house 22'")

	http.ListenAndServe(":8080", nil)
}

func emergencyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the body
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	rawMessage := string(body)

	// Parse it (reuse our logic from before)
	msg := ParseMsg(rawMessage)

	fmt.Println("NEW EMERGENCY:")
	fmt.Printf("House: %s\n", msg.HouseNumber)
	fmt.Printf("Priority: %s\n", msg.Priority)
	fmt.Printf("Source: %s\n------\n", msg.Source)
	if msg.Injury != "" {
		fmt.Printf("Injury: %s\n------\n", msg.Injury)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "received", "house": "%s", "priority": "%s"}`, msg.HouseNumber, msg.Priority)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
