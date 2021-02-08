package goml

// Regression includes samples, targets and prediction outcome
type Regression struct {
	Samples [][]float64
	Targets [][]float64
	Predict []float64
}

// Train datasets
func (e *Regression) Train(samples [][]float64, targets []float64) {
	// merge samples
	for k, v := range samples {
		e.Samples[k] = append(e.Samples[k], v...)
	}

	// merge targets
	for k := range samples {
		e.Targets[k] = append(e.Targets[k], targets...)
	}

	e.computeCoefficients()
}

// func (e *Regression) Predict(samples [][]float64) float64 {

// }

func predictSample() {

}

// coefficient(b) = (X'X)-1X'Y
func (e *Regression) computeCoefficients() error {
	samplesMatrix := e.getSamplesMatrix()
	targetsMatrix := e.getTargetsMatrix()

	ts, err := samplesMatrix.transpose().multiply(samplesMatrix.samples)
	tf, er := samplesMatrix.transpose().multiply(targetsMatrix.samples)

	if err != nil {
		return err
	}

	ts.inverse()

	if er != nil {
		return er
	}

	ts.multiply(tf.samples)

	return nil
}

func (e *Regression) getSamplesMatrix() *matrix {
	var samples [][]float64

	for k, v := range e.Samples {
		samples[k] = append([]float64{1}, v...)
	}

	return &matrix{samples: samples}
}

func (e *Regression) getTargetsMatrix() *matrix {
	return &matrix{samples: e.Targets}
}
