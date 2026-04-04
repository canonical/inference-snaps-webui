package ui

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Serve(conf Config, htmlDir string, port int, bindAddress string) error {

	// Error out if httpDir doesn't contain an index.html file, as that's required for the frontend to work
	if _, err := http.Dir(htmlDir).Open("index.html"); err != nil {
		return fmt.Errorf("Failed to find index.html in %q: %w", htmlDir, err)
	}

	mux := http.NewServeMux()

	// Serve configuration for the frontend
	mux.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-store")
		if err := json.NewEncoder(w).Encode(conf); err != nil {
			http.Error(w, "failed to encode config", http.StatusInternalServerError)
		}
	})

	// Serve the frontend static files
	mux.Handle("/", http.FileServer(http.Dir(htmlDir)))

	fmt.Printf("Serving %q on http://localhost:%d\n", htmlDir, port)

	return http.ListenAndServe(fmt.Sprintf("%s:%d", bindAddress, port), mux)
}
