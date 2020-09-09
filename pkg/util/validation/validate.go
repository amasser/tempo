package validation

import "github.com/prometheus/client_golang/prometheus"

const (
	discardReasonLabel = "reason"

	// RateLimited is one of the values for the reason to discard samples.
	// Declared here to avoid duplication in ingester and distributor.
	RateLimited = "rate_limited"
)

// DiscardedBytes is a metric of the total discarded bytes, by reason.
var DiscardedBytes = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "tempo",
		Name:      "discarded_bytes_total",
		Help:      "The total number of bytes that were discarded.",
	},
	[]string{discardReasonLabel, "tenant"},
)

// DiscardedSamples is a metric of the number of discarded samples, by reason.
var DiscardedSamples = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "tempo",
		Name:      "discarded_samples_total",
		Help:      "The total number of samples that were discarded.",
	},
	[]string{discardReasonLabel, "tenant"},
)

func init() {
	prometheus.MustRegister(DiscardedSamples, DiscardedBytes)
}

// ValidTraceID confirms that trace ids are 128 bits
func ValidTraceID(id []byte) bool {
	return len(id) == 16
}