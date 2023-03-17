package entity

import "testing"

func TestGetAllMetrics(t *testing.T) {
	allMetrics := GetAll()

	if len(allMetrics) != len(MetricsList) {
		t.Fail()
	}
	for i, v := range allMetrics {
		if v.Type != allMetrics[i].Type ||
			v.Name != allMetrics[i].Name ||
			v.Value != allMetrics[i].Value ||
			v.Id != allMetrics[i].Id {
			t.Fail()
			break
		}
	}
}
