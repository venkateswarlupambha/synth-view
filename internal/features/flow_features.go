package features

import (
	"math"
	"synthview/internal/flow"
)


// FlowFeatures represents extracted statistical features for a flow
type FlowFeatures struct {
	SrcIP       string
	DstIP       string
	Protocol    string
	Duration    float64
	PacketCount int
	ByteCount   int
	PacketRate  float64
	ByteRate    float64
	IATMean     float64
	IATVariance float64
	Burstiness  float64
}

// FromFlowStats extracts features from FlowStats
func FromFlowStats(f *flow.FlowStats) FlowFeatures {

	duration := f.LastSeen.Sub(f.StartTime).Seconds()
	if duration <= 0 {
		duration = 0.000001 // avoid division by zero
	}

	packetRate := float64(f.PacketCount) / duration
	byteRate := float64(f.ByteCount) / duration

	iatMean, iatVar := computeIATStats(f.IATs)
	burstiness := 0.0
	if iatMean > 0 {
burstiness = math.Sqrt(iatVar) / iatMean

	}

	return FlowFeatures{
		SrcIP:       f.Key.SrcIP,
		DstIP:       f.Key.DstIP,
		Protocol:    f.Key.Protocol,
		Duration:    duration,
		PacketCount: f.PacketCount,
		ByteCount:   f.ByteCount,
		PacketRate:  packetRate,
		ByteRate:    byteRate,
		IATMean:     iatMean,
		IATVariance: iatVar,
		Burstiness:  burstiness,
	}
}

