package aclreceiver

import "go.opentelemetry.io/collector/component"

// define whatever that you want to set in the config.yaml
type CustomReceiverConfig struct{}

func createDefaultConfig() component.Config {
	return &CustomReceiverConfig{}
}
