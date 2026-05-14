package baseline

import "synthview/internal/features"

// BuildBaseline computes protocol-wise mean feature values
func BuildBaseline(featureList []features.FlowFeatures) BaselineModel {

	type accum struct {
		count     int
		duration  float64
		pktRate   float64
		byteRate  float64
		iatVar    float64
		burst     float64
	}

	temp := make(map[string]*accum)

	for _, f := range featureList {

		if _, ok := temp[f.Protocol]; !ok {
			temp[f.Protocol] = &accum{}
		}

		a := temp[f.Protocol]
		a.count++
		a.duration += f.Duration
		a.pktRate += f.PacketRate
		a.byteRate += f.ByteRate
		a.iatVar += f.IATVariance
		a.burst += f.Burstiness
	}

	model := BaselineModel{
		Profiles: make(map[string]BaselineProfile),
	}

	for proto, a := range temp {

		model.Profiles[proto] = BaselineProfile{
			Protocol:     proto,
			MeanDuration: a.duration / float64(a.count),
			MeanPktRate:  a.pktRate / float64(a.count),
			MeanByteRate: a.byteRate / float64(a.count),
			MeanIATVar:   a.iatVar / float64(a.count),
			MeanBurst:    a.burst / float64(a.count),
		}
	}

	return model
}

