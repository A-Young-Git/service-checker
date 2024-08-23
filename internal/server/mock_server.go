/*mock_server.go implements a mock HTTP server to simulate
running services for testing purposes.*/
package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// StartMockService starts a mock server for the specified service
func StartMockService() {
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("SERVICE_NAME environment variable is required")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s is running\n", serviceName)
	})
	log.Printf("%s is running on port 8090", serviceName)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
