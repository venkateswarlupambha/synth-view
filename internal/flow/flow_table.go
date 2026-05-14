package flow

import (
	"sync"
	

	"synthview/internal/capture"
)

// FlowTable manages active flows
type FlowTable struct {
	flows map[FlowKey]*FlowStats
	mu    sync.Mutex
}

// NewFlowTable initializes a flow table
func NewFlowTable() *FlowTable {
	return &FlowTable{
		flows: make(map[FlowKey]*FlowStats),
	}
}

// ProcessPacket updates flow statistics using a PacketEvent
func (ft *FlowTable) ProcessPacket(pkt capture.PacketEvent) {

	ft.mu.Lock()
	defer ft.mu.Unlock()

	key := FlowKey{
		SrcIP:    pkt.SrcIP,
		DstIP:    pkt.DstIP,
		SrcPort:  pkt.SrcPort,
		DstPort:  pkt.DstPort,
		Protocol: pkt.Protocol,
	}

now := pkt.Timestamp


	flow, exists := ft.flows[key]

	if !exists {
		ft.flows[key] = &FlowStats{
			Key:          key,
			StartTime:    now,
			LastSeen:     now,
			PacketCount: 1,
			ByteCount:   pkt.Length,
			IATs:         []float64{},
		}
		return
	}

	iat := now.Sub(flow.LastSeen).Seconds()
	flow.IATs = append(flow.IATs, iat)

	flow.PacketCount++
	flow.ByteCount += pkt.Length
	flow.LastSeen = now
}

// GetFlows returns all current flows
func (ft *FlowTable) GetFlows() []*FlowStats {

	ft.mu.Lock()
	defer ft.mu.Unlock()

	var result []*FlowStats
	for _, flow := range ft.flows {
		result = append(result, flow)
	}
	return result
}
