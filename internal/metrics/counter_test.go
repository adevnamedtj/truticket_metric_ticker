package metrics

import (
	"math/rand"
	"testing"
)

func TestUpdateMetric(t *testing.T) {

	value := rand.Float64()
	type args struct {
		metricType  string
		metricName  string
		metricCount float64
	}
	tests := []struct {
		name      string
		args      args
		wantCount float64
	}{
		{name: "add and get metric test", args: args{metricType: "counter", metricName: "test", metricCount: value}, wantCount: value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if gotCount := UpdateMetric(tt.args.metricType, tt.args.metricName, tt.args.metricCount); gotCount != tt.wantCount {
				t.Errorf("GetCurrentTimeBucketMetric() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
