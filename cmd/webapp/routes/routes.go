package routes

import (
	"encoding/json"
	"github.com/ckalagara/truticket_metric_ticker/internal/metrics"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"math"
	"net/http"
	"strings"
)

func GetRoutes() chi.Router {
	logrus.Info("Initializing the routes...")
	r := chi.NewRouter()

	r.Get("/{key}/sum", GetCurrentTimeBucketMetricHandler)
	r.Post("/{key}", UpdateMetricHandler)

	return r
}

type Payload struct {
	Value interface{} `json:"value"`
}

// GetCurrentTimeBucketMetricHandler handle get metric request responds with sum of reported values in current time bucket
func GetCurrentTimeBucketMetricHandler(w http.ResponseWriter, r *http.Request) {
	metricKey := strings.TrimSpace(chi.URLParam(r, "key"))
	if len(metricKey) == 0 {
		http.Error(w, "Please provide a valid metric as part of path, eg: /metric/{key}.", http.StatusBadRequest)
	}

	val := metrics.GetCurrentTimeBucketMetric("counter", metricKey)
	w.Header().Set("Content-Type", "application/json")
	writePayload(w, Payload{Value: math.Round(val)})

}

// UpdateMetricHandler handle request to post a metric value
func UpdateMetricHandler(w http.ResponseWriter, r *http.Request) {

	metricKey := strings.TrimSpace(chi.URLParam(r, "key"))
	if len(metricKey) == 0 {
		http.Error(w, "Please provide a valid metric as part of path, eg: /metric/{key}/sum.", http.StatusBadRequest)
	}
	var p Payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":   err,
			"payload": p,
		}).Errorf("failed to json decode Payload from request")
		http.Error(w, "Please provide a valid request body.", http.StatusBadRequest)
	}

	_ = metrics.UpdateMetric("counter", metricKey, p.Value.(float64))
	writePayload(w, Payload{})

}

func writePayload(w http.ResponseWriter, p Payload) {
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":   err,
			"payload": p,
		}).Errorf("failed to json encode Payload to response writer")
		http.Error(w, "Something went wrong, please try later.", http.StatusInternalServerError)
	}
}
