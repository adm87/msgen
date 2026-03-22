package models

type MSGenCfg struct {
	Service ServiceCfg `yaml:"service"`
}

type ServiceCfg struct {
	Name   string `yaml:"name"`
	Module string `yaml:"module"`
}
