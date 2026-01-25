package models

const MSConfigFileName = "msgen.json"

type MSGenConfig struct {
	CodegenVersion string `json:"codegen_version,omitempty"`
	ModuleUrl      string `json:"module_url,omitempty"`
}

func DefaultMSGenConfig() MSGenConfig {
	return MSGenConfig{
		CodegenVersion: "1.0.0",
		ModuleUrl:      "github.com/yourusername/yourmodule",
	}
}
