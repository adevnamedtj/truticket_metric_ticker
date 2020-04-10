package metrics

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var metricData sync.Map

var timeBucket = os.Getenv("APP_METRICS_TIME_BUCKET_INTERVAL")
var timeKeyLayout string

func init() {

	switch timeBucket {

	case "minutes":
		timeKeyLayout = "2006-01-02 15:04"

	case "hours":
		timeKeyLayout = "2006-01-02 15"

	case "days":
		timeKeyLayout = "2006-01-02"

	default:
		timeKeyLayout = "2006-01-02 15"

	}

}

// UpdateMetric takes a metric type, name and value and create/update the current time bucket(configured based on timeKeyLayout/timeBucket keys) count for the metric and returns the current count
func UpdateMetric(metricType, metricName string, metricCount float64) (count float64) {

	key := fmt.Sprintf("%s_%s_%s", metricType, metricName, getCurrentTimeBucketKey())
	currentCounter, exists := metricData.Load(key)
	if exists {
		metricData.Store(key, currentCounter.(float64)+metricCount)
	} else {
		metricData.Store(key, metricCount)
	}

	return GetCurrentTimeBucketMetric(metricType, metricName)
}

// GetCurrentTimeBucketMetric takes in metric type and name and returns the count of current time bucket value (configured based on timeKeyLayout/timeBucket keys)
func GetCurrentTimeBucketMetric(metricType, metricName string) (count float64) {
	key := fmt.Sprintf("%s_%s_%s", metricType, metricName, getCurrentTimeBucketKey())
	currentCounter, exists := metricData.Load(key)
	if exists {
		count = currentCounter.(float64)
	}
	return
}

func getCurrentTimeBucketKey() (k string) {
	//formatting current time to identify up to hour indicator, to return time in hourly keys
	k = time.Now().Format(timeKeyLayout)
	return
}
