package prometheus

type Metric struct {
	Metric string
	Labels map[string]string
	Value  int64
}
