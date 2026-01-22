package main

import (
	"encoding/json"
	"log"
	"net/http"

	"rpi-status-dashboard-go/stats"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		data, err := stats.Get()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	mux.Handle("/", http.FileServer(http.Dir("./web")))

	log.Println("🟢 RPi status dashboard running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
