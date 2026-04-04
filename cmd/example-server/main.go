package main

import (
	"log"

	ui "github.com/canonical/inference-snaps-ui"
)

const staticDir = "static"

func main() {
	config := ui.Config{
		OpenAIBaseURL: "http://localhost:8328/v1",
		Capabilities:  ui.SupportedUICapabilities(),
		InstanceName:  "example",
		EngineName:    "intel-cpu",
	}

	if err := ui.Serve(config, staticDir, 8080, "localhost"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
