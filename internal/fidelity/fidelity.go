package fidelity

import (
	"math"

	"synthview/internal/baseline"
	"synthview/internal/features"
)

// ComputeFidelity calculates a 0–100 fidelity score for a flow
func ComputeFidelity(
	flow features.FlowFeatures,
	base baseline.BaselineProfile,
) float64 {

	const eps = 1e-6

	// Normalized differences (scale-invariant)
	d1 := (flow.Duration - base.MeanDuration) / (base.MeanDuration + eps)
	d2 := (flow.PacketRate - base.MeanPktRate) / (base.MeanPktRate + eps)
	d3 := (flow.Burstiness - base.MeanBurst) / (base.MeanBurst + eps)

	// Euclidean distance in normalized space
	distance := math.Sqrt(d1*d1 + d2*d2 + d3*d3)

	// Convert distance to fidelity score
	score := 100.0 * math.Exp(-distance)

	if score < 0 {
		return 0
	}
	if score > 100 {
		return 100
	}

	return score
}


