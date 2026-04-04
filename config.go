package ui

type Config struct {
	OpenAIBaseURL string   `json:"openAIBaseURL"`
	Capabilities  []string `json:"capabilities"`
	InstanceName  string   `json:"instanceName"`
	EngineName    string   `json:"engineName"`
}

const (
	UICapabilityText   string = "text"
	UICapabilityVision string = "vision"
)

func SupportedUICapabilities() []string {
	return []string{UICapabilityText, UICapabilityVision}
}
