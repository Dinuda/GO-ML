package goml

type Estimator interface {
	Train([][]float64, [][]float64)
	Predict() float64
}
