# ☸️ KubeTask Management System

A cloud-native Go application deployed on Kubernetes with Helm-based monitoring using Prometheus and Grafana.

## Tech Stack

- Go (Golang)
- Docker
- Kubernetes
- Helm
- Prometheus
- Grafana

## Architecture

Go Application
|
|
Docker Container
|
|
Kubernetes Deployment
|
|
Service (NodePort)
|
|
Prometheus → Grafana


## Features

Go REST API service  
Docker containerized application  
Kubernetes Deployment with replicas  
Kubernetes Service exposure  
ConfigMap configuration  
Health endpoint  
Prometheus monitoring  
Grafana dashboards


## Kubernetes

Deployment:

- 2 replicas
- rolling updates
- pod management


Service:

- NodePort exposure
- internal Kubernetes networking


## Monitoring

Implemented:

- Prometheus metrics collection
- Grafana visualization

Dashboard includes:

- Request rate
- Latency monitoring
- Error rate
- Pod health
- Resource usage


## Run Locally

### Build Docker image

```bash
docker build -t kubetask .