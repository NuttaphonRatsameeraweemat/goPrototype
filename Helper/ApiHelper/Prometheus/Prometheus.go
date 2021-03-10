package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

// SetGaugeMetric func(name string, help string, env string, envValue string, version string, versionValue string) (prometheusGauge Gauge)
func SetGaugeMetric(name string, help string, env string, envValue string, version string, versionValue string) (prometheusGauge prometheus.Gauge) {
	var (
		gaugeMetric = prometheus.NewGauge(prometheus.GaugeOpts{
			Name:        name,
			Help:        help,
			ConstLabels: prometheus.Labels{env: envValue, version: versionValue},
		})
	)

	return gaugeMetric
}
