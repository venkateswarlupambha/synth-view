package main
import (
	"fmt"

	"synthview/internal/capture"
	"synthview/internal/features"
	"synthview/internal/flow"
	"synthview/internal/baseline"
	"synthview/internal/fidelity"
	"synthview/internal/fingerprint"
	"synthview/internal/metrics"
	"synthview/internal/api"


)


func main() {

	packetChannel := make(chan capture.PacketEvent)

	flowTable := flow.NewFlowTable()

	go capture.ReadPCAP("datasets/sample.pcap", packetChannel)

	for pkt := range packetChannel {
		flowTable.ProcessPacket(pkt)
	}

	fmt.Println("\n=== FLOW FEATURES ===")

for _, f := range flowTable.GetFlows() {

	feat := features.FromFlowStats(f)

	fmt.Printf(
		"%s → %s | %s | dur=%.3fs pkts=%d bytes=%d pkt_rate=%.2fB/s iat_var=%.6f burst=%.3f\n",
		feat.SrcIP,
		feat.DstIP,
		feat.Protocol,
		feat.Duration,
		feat.PacketCount,
		feat.ByteCount,
		feat.PacketRate,
		feat.IATVariance,
		feat.Burstiness,
	)
}
var featureList []features.FlowFeatures

for _, f := range flowTable.GetFlows() {
	feat := features.FromFlowStats(f)
	featureList = append(featureList, feat)
}

baselineModel := baseline.BuildBaseline(featureList)

fmt.Println("\n=== BASELINE MODEL ===")
for proto, b := range baselineModel.Profiles {
	fmt.Printf(
		"%s | mean_dur=%.4f mean_pkt_rate=%.2f mean_burst=%.3f\n",
		proto,
		b.MeanDuration,
		b.MeanPktRate,
		b.MeanBurst,
	)
}
fmt.Println("\n=== FIDELITY SCORES ===")

for _, feat := range featureList {

	base, ok := baselineModel.Profiles[feat.Protocol]
	if !ok {
		continue
	}

	score := fidelity.ComputeFidelity(feat, base)

	fmt.Printf(
		"%s → %s | %s | Fidelity Score = %.2f\n",
		feat.SrcIP,
		feat.DstIP,
		feat.Protocol,
		score,
	)
}
fmt.Println("\n=== FLOW FINGERPRINTING ===")

for _, feat := range featureList {

	base, ok := baselineModel.Profiles[feat.Protocol]
	if !ok {
		continue
	}

	score := fidelity.ComputeFidelity(feat, base)
	label := fingerprint.ClassifyFlow(feat, score)

	fmt.Printf(
		"%s → %s | %s | Fidelity=%.2f | Fingerprint=%s\n",
		feat.SrcIP,
		feat.DstIP,
		feat.Protocol,
		score,
		label,
	)
}
var fidelities []float64
var labels []string

for _, feat := range featureList {

	base, ok := baselineModel.Profiles[feat.Protocol]
	if !ok {
		continue
	}

	score := fidelity.ComputeFidelity(feat, base)
	label := fingerprint.ClassifyFlow(feat, score)

	fidelities = append(fidelities, score)
	labels = append(labels, string(label))
}

	sysMetrics := metrics.ComputeMetrics(fidelities, labels)

	fmt.Println("\n=== SYSTEM METRICS ===")
	fmt.Printf("Total Flows: %d\n", sysMetrics.TotalFlows)
	fmt.Printf("Average Fidelity: %.2f\n", sysMetrics.AvgFidelity)
	fmt.Printf("Low-Fidelity Flows (<60): %d\n", sysMetrics.LowFidelity)
	fmt.Printf("DNS-like Flows: %d\n", sysMetrics.DNSFlows)
	fmt.Printf("HTTP-like Flows: %d\n", sysMetrics.HTTPFlows)

go api.StartServer(sysMetrics)

fmt.Println("\n[✔] API server running at http://localhost:8080/metrics")

select {} // keep program running



}

