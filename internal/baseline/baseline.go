package baseline

// BaselineProfile represents averaged feature values for a protocol
type BaselineProfile struct {
	Protocol     string
	MeanDuration float64
	MeanPktRate  float64
	MeanByteRate float64
	MeanIATVar   float64
	MeanBurst    float64
}

// BaselineModel stores baselines per protocol
type BaselineModel struct {
	Profiles map[string]BaselineProfile
}

