package main

import (
	"github.com/ckalagara/truticket_metric_ticker/cmd/webapp/routes"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var port = os.Getenv("APP_PORT")

func init() {
	if len(port) == 0 {
		port = ":3333"
	}
}

func main() {

	logrus.Info("Initializing the service...")
	r := chi.NewRouter()

	r.Mount("/metric", routes.GetRoutes())

	logrus.Infof("Initializing http server over port %v...", port)
	err := http.ListenAndServe(port, r)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Errorf("Failed to create http server")
	}
}
