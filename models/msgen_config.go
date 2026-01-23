package models

const MSConfigFileName = "msgen.json"

type MSGenConfig struct {
	CodegenVersion string `json:"codegen_version,omitempty"`
	ServiceName    string `json:"service_name,omitempty"`
	Port           int    `json:"port,omitempty"`
}

func DefaultMSGenConfig() MSGenConfig {
	return MSGenConfig{
		CodegenVersion: "1.0.0",
		ServiceName:    "new-service",
		Port:           8080,
	}
}
