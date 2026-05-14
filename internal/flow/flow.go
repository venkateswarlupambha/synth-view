package flow

import "time"

// FlowKey uniquely identifies a network flow (5-tuple)
type FlowKey struct {
	SrcIP    string
	DstIP    string
	SrcPort  string
	DstPort  string
	Protocol string
}

// FlowStats stores aggregated statistics for a flow
type FlowStats struct {
	Key          FlowKey
	StartTime    time.Time
	LastSeen     time.Time
	PacketCount int
	ByteCount   int
	IATs         []float64
}
