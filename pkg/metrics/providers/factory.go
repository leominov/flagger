package providers

import (
	"go.uber.org/zap"

	flaggerv1 "github.com/weaveworks/flagger/pkg/apis/flagger/v1beta1"
)

type Factory struct{}

func (factory Factory) Provider(
	logger *zap.SugaredLogger,
	metricInterval string,
	provider flaggerv1.MetricTemplateProvider,
	credentials map[string][]byte,
) (Interface, error) {
	switch provider.Type {
	case "prometheus":
		return NewPrometheusProvider(logger, provider, credentials)
	case "datadog":
		return NewDatadogProvider(logger, metricInterval, provider, credentials)
	case "cloudwatch":
		return NewCloudWatchProvider(logger, metricInterval, provider)
	default:
		return NewPrometheusProvider(logger, provider, credentials)
	}
}
