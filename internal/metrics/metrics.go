package metrics

type Metrics struct {
	TotalFlows     int
	AvgFidelity    float64
	LowFidelity    int
	DNSFlows       int
	HTTPFlows      int
}

// ComputeMetrics aggregates system-level metrics
func ComputeMetrics(fidelities []float64, labels []string) Metrics {

	var sum float64
	low := 0
	dns := 0
	http := 0

	for i, f := range fidelities {
		sum += f
		if f < 60 {
			low++
		}

		switch labels[i] {
		case "DNS-like":
			dns++
		case "HTTP-like":
			http++
		}
	}

	avg := 0.0
	if len(fidelities) > 0 {
		avg = sum / float64(len(fidelities))
	}

	return Metrics{
		TotalFlows:  len(fidelities),
		AvgFidelity: avg,
		LowFidelity: low,
		DNSFlows:    dns,
		HTTPFlows:  http,
	}
}

