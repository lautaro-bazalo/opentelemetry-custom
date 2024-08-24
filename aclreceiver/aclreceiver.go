package aclreceiver

import (
	"context"
	"log"
	"strings"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type ACLReceiver struct {
	consumer.Metrics
	nextConsumer consumer.Metrics
	Config       *CustomReceiverConfig
	host         component.Host
	cancel       context.CancelFunc
	logger       *zap.Logger
}

func (r *ACLReceiver) Start(ctx context.Context, host component.Host) error {
	// Aquí inicializas la conexión con el stream de logs o cualquier otro setup
	r.host = host
	ctx = context.Background()
	ctx, r.cancel = context.WithCancel(ctx)
	go r.consumeLogs(ctx)
	return nil
}

func (r *ACLReceiver) consumeLogs(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			// I simulate receive a log
			logData := "GET /api/v1/resource 200"

			// Parsing logs
			parts := strings.Split(logData, " ")
			if len(parts) < 3 {
				r.logger.Sugar().Errorf("Log malformed: %s", logData)
				continue
			}
			method := parts[0]
			resource := parts[1]
			statusCode := parts[2]

			// Generating metrics
			metrics := pmetric.NewMetrics()
			rm := metrics.ResourceMetrics().AppendEmpty()
			sm := rm.ScopeMetrics().AppendEmpty()
			metric := sm.Metrics().AppendEmpty()
			metric.SetName("http_requests_total")
			metric.SetUnit("1")
			sum := metric.SetEmptySum()
			sum.SetIsMonotonic(true)
			sum.SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
			dp := sum.DataPoints().AppendEmpty()
			dp.SetIntValue(1)
			dp.Attributes().PutStr("method", method)
			dp.Attributes().PutStr("resource", resource)
			dp.Attributes().PutStr("status_code", statusCode)
			dp.SetTimestamp(pmetric.NewHistogram().DataPoints().AppendEmpty().StartTimestamp())

			err := r.nextConsumer.ConsumeMetrics(ctx, metrics)
			if err != nil {
				log.Printf("Error sending the metrics: %v", err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (r *ACLReceiver) Shutdown(ctx context.Context) error {
	return nil
}
