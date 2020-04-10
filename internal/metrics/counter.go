package metrics

import (
	"fmt"
	"sync"
	"time"
)

var metricData sync.Map

// UpdateMetric takes a metric type, name and value and create/update the current hour bucket count for the metric and returns the current count
func UpdateMetric(metricType, metricName string, metricCount float64) (count float64) {

	key := fmt.Sprintf("%s_%s_%s", metricType, metricName, getCurrentTimeBucketKey())
	currentCounter, exists := metricData.Load(key)
	if exists {
		metricData.Store(key, currentCounter.(float64)+metricCount)
	} else {
		metricData.Store(key, metricCount)
	}

	return GetMetric(metricType, metricName)
}

// GetMetric takes in metric type and name and returns the count of current hour bucket value
func GetMetric(metricType, metricName string) (count float64) {
	key := fmt.Sprintf("%s_%s_%s", metricType, metricName, getCurrentTimeBucketKey())
	currentCounter, exists := metricData.Load(key)
	if exists {
		count = currentCounter.(float64)
	}
	return
}

func getCurrentTimeBucketKey() (k string) {
	//formatting current time to identify up to hour indicator, to return time in hourly keys
	k = time.Now().Format("2006-01-02 15")
	return
}
