// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"io"
	"net/http"
	"os"

	"github.com/devcastops/gcp_server_control/modules/instances"
	"github.com/google/logger"
)

func main() {
	logger := logger.Init("main log", true, true, io.Discard)
	defer logger.Close()
	logger.Info("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		logger.Infof("defaulting to port %s", port)
	}

	// Start HTTP server.
	logger.Infof("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	logger := logger.Init("hander log", true, true, w)
	defer logger.Close()
	logger.Info(instances.ListAllInstances(logger))
}
