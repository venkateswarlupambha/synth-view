package fingerprint

import "synthview/internal/features"

// FingerprintLabel represents classified traffic type
type FingerprintLabel string

const (
	DNSLike     FingerprintLabel = "DNS-like"
	HTTPLike    FingerprintLabel = "HTTP-like"
	Anomalous   FingerprintLabel = "Anomalous"
)

// ClassifyFlow assigns a fingerprint label based on flow features
func ClassifyFlow(f features.FlowFeatures, fidelity float64) FingerprintLabel {

	// DNS fingerprint (UDP, single-packet)
	if f.Protocol == "UDP" && f.PacketCount <= 2 {
		return DNSLike
	}

	// HTTP fingerprint (TCP, bursty, short-lived)
	if f.Protocol == "TCP" &&
		f.PacketCount >= 4 &&
		f.Burstiness >= 0.5 &&
		f.Duration > 0 {
		return HTTPLike
	}

	// Anything else is anomalous
	if fidelity < 60 {
		return Anomalous
	}

	return Anomalous
}

