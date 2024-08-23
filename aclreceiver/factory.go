package aclreceiver

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var ACLReceiverType = component.MustNewType("aclreceiver")

func createACLReceiver(_ context.Context, params receiver.Settings, baseConfig component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {

	l := params.Logger

	return &ACLReceiver{
		nextConsumer: consumer,
		Config:       baseConfig.(*CustomReceiverConfig),
		logger:       l,
	}, nil
}

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		ACLReceiverType,
		createDefaultConfig,
		receiver.WithMetrics(createACLReceiver, component.StabilityLevelBeta))
}
