package metric

import (
	"errors"
)

const ErrInvalidType = "invalid type of metric"
const ErrUnknownMetric = "unknown metric"

type Metrics struct {
	gauge   map[string]string
	counter map[string]string
	test    int
}

func NewMetrics() *Metrics {
	gaugeMetrics := make(map[string]string)
	counterMetrics := make(map[string]string)
	return &Metrics{gauge: gaugeMetrics, counter: counterMetrics, test: 1}
}

func (m Metrics) MetricsOfType(typeM string) (map[string]string, error) {
	switch typeM {
	case "gauge":
		return m.gauge, nil
	case "counter":
		return m.counter, nil
	default:
		return nil, errors.New(ErrInvalidType)
	}
}
