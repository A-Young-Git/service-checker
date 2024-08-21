# service-checker
Service Status Checker CLI Tool

This project is part of a series of 1-2 hour projects designed to sharpen Go fundamentals with a focus on DevOps practices.

## Project Overview
This project demonstrates how to create a simple CLI tool in Go that checks the status of services running on multiple servers. The tool is designed to run service checks in parallel, improving efficiency when dealing with multiple servers. The steps outlined below guide you through the process of structuring the application using best practices, writing the core functionality, and simulating service checks locally.

## Steps

### Step 1: Create Mock Servers
Simulate the behavior of different services (e.g., `nginx`, `docker`) locally by creating mock servers in Go. These servers will respond with a "running" message when queried.

### Step 2: Implement the Service Status Checker CLI
Develop the core functionality of the CLI tool in Go. This tool will:
- Accept a list of servers and services to check.
- Perform service status checks in parallel using goroutines.
- Output the status of each service on each server.

### Step 3: Structure the Project
Organize the code into multiple packages:
- `server`: Manages server configurations and mock server setup.
- `service`: Handles service status checks and responses.
- `main`: The entry point for the application, coordinating the CLI and service checks.

### Step 4: Run and Test the Application
With the mock servers running, execute the service checker to verify that it correctly identifies the status of services. Test different scenarios, such as a service being down, and observe how the tool responds.

### Step 5: Extend and Optimize
Consider additional features such as integrating real SSH connections, adding support for more services, or outputting the results in different formats (e.g., JSON).

---

This project reinforces Go fundamentals, parallel processing, and DevOps principles, preparing you for more advanced scenarios in system administration and automation.
