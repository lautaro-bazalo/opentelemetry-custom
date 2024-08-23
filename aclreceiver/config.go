package aclreceiver

import "go.opentelemetry.io/collector/component"

type CustomReceiverConfig struct{}

func createDefaultConfig() component.Config {
	return &CustomReceiverConfig{}
}
