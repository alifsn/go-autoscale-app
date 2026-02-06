# Go Autoscale App

A simple Go web application designed to demonstrate Kubernetes Horizontal Pod Autoscaling (HPA) capabilities.

## Overview

This application creates a CPU-intensive web server that simulates load to trigger Kubernetes autoscaling. It's built to showcase how Kubernetes can automatically scale pods based on CPU utilization.

## Features

- **CPU Load Simulation**: Performs mathematical calculations to generate CPU load
- **Pod Identification**: Returns the hostname (pod name) in responses
- **Kubernetes Ready**: Configured with proper resource limits for HPA
- **Containerized**: Docker image ready for deployment

## Repository Structure

```
go-autoscale-app/
├── main.go           # Main Go application
├── go.mod            # Go module definition
├── Dockerfile        # Container build configuration
├── deployment.yaml   # Kubernetes deployment manifest
├── service.yaml      # Kubernetes service configuration
├── hpa.yaml          # Horizontal Pod Autoscaler configuration
└── README.md         # This file
```

## Application Details

### Main Application (`main.go`)
- HTTP server listening on port 8080
- Single endpoint `/` that:
  - Performs CPU-intensive calculations (10M iterations)
  - Returns pod hostname and calculation result
  - Simulates realistic CPU load for autoscaling

### Docker Configuration (`Dockerfile`)
- Based on `golang:1.24-alpine`
- Builds static binary
- Exposes port 8080
- Lightweight Alpine Linux base

## Kubernetes Configuration

### Deployment (`deployment.yaml`)
- **Initial replicas**: 1 pod
- **Resource requests**: 100m CPU (0.1 cores)
- **Resource limits**: 500m CPU (0.5 cores)
- **Container port**: 8080

### Service (`service.yaml`)
- **Type**: ClusterIP
- **Port mapping**: 80 → 8080
- **Selector**: `app: go-autoscale`

### HPA (`hpa.yaml`)
- **Min replicas**: 1
- **Max replicas**: 5
- **CPU threshold**: 50% utilization
- **Target**: go-autoscale deployment

## Deployment Instructions

### Prerequisites
- Kubernetes cluster
- kubectl configured
- Metrics server installed

### Deploy Application
```bash
# Apply all Kubernetes manifests
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f hpa.yaml

# Verify deployment
kubectl get pods
kubectl get hpa
```

### Test Autoscaling
```bash
# Generate load to trigger scaling
kubectl run -i --tty load-generator --rm --image=busybox --restart=Never -- /bin/sh

# Inside the pod, run:
while true; do wget -q -O- http://go-autoscale-service/; done
```

### Monitor Scaling
```bash
# Watch HPA status
kubectl get hpa go-autoscale-hpa --watch

# Monitor pods
kubectl get pods --watch
```

## Building Docker Image

```bash
# Build image
docker build -t alifsn/go-autoscale:latest .

# Push to registry
docker push alifsn/go-autoscale:latest
```

## Expected Behavior

1. **Normal Load**: Application runs with 1 replica
2. **High Load**: When CPU usage exceeds 50%, HPA scales up to max 5 replicas
3. **Load Decrease**: When load drops, HPA scales down after cooldown period

## Use Cases

- **Learning Kubernetes HPA**: Understand autoscaling concepts
- **Load Testing**: Test cluster autoscaling behavior
- **Demo Purposes**: Showcase Kubernetes capabilities
- **Development**: Base for building scalable applications

## Requirements

- Go 1.24+
- Kubernetes 1.20+
- Metrics Server enabled
- Docker (for containerization)