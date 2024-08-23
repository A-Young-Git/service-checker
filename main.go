/*
This main.go file sets up a list of servers and services, then uses
goroutines to check the status of each service in parallel. The results are
collected and printed after all checks are complete.
*/
package main

import (
	"fmt"
	"service-checker/internal/server"
	"service-checker/internal/service"
	"strings"
	"sync"
)

func main() {
	go server.StartMockService()
	servers := []server.Server{
		{IP: "localhost", Services: []string{"nginx"}},
		{IP: "localhost", Services: []string{"docker"}},
	}

	var wg sync.WaitGroup
	results := make(chan string, len(servers)*len(servers[0].Services))

	for _, srv := range servers {
		for _, svc := range srv.Services {
			wg.Add(1)
			go service.CheckServiceStatus(srv.IP, strings.TrimSpace(svc), &wg, results)
		}
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}

	select {}
}
