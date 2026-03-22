package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// config holds the deployment configuration.
// Edit these values to customise the application for a specific deployment.
var config = struct {
	OpenAIBaseURL  string `json:"openAIBaseURL"`
	SupportsImages bool   `json:"supportsImages"`
	ReasoningModel bool   `json:"reasoningModel"`
	UIName         string `json:"uiName"`
	EngineName     string `json:"engineName"`
}{
	OpenAIBaseURL:  "http://localhost:8328/v3",
	SupportsImages: true,
	ReasoningModel: false,
	UIName:         "Gemma3 Inference Snap",
	EngineName:     "intel-cpu",
}

func main() {
	mux := http.NewServeMux()

	// Serve configuration as JSON from the Go config values above.
	mux.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-store")
		if err := json.NewEncoder(w).Encode(config); err != nil {
			http.Error(w, "failed to encode config", http.StatusInternalServerError)
		}
	})

	// Serve the frontend static files from the webui/ directory.
	mux.Handle("/", http.FileServer(http.Dir("webui")))

	addr := ":8080"
	log.Printf("Serving %s on http://localhost%s", config.UIName, addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
