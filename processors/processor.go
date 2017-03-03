package processors

import (
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/alphagov/paas-cf-apps-statsd/metrics"
)

type Processor interface {
	Process(e *events.Envelope) ([]metrics.Metric, error)
}
