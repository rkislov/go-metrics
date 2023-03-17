package entity

type Metric struct {
	Type  string
	Name  string
	Value float64
}

var MetricsList = []Metric{
	{Type: "gauge", Name: "Alloc", Value: 0000},
	{Type: "gauge", Name: "BuckHashSys", Value: 0001},
}
