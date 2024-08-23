/* service.go checks if the specificed services are running on the specified server IP.*/
package service

import (
	"fmt"
	"net/http"
	"sync"
)

// CheckServiceStatus checks if a service is running on the given server IP
func CheckServiceStatus(ip string, serviceName string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	url := fmt.Sprintf("http://%s:8090", ip)
	resp, err := http.Get(url)
	status := "Not Running"

	if err == nil && resp.StatusCode == http.StatusOK {
		status = "Running"
	}

	results <- fmt.Sprintf("Server: %s - Service: %s - Status: %s", ip, serviceName, status)
}
