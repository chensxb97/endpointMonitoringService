## Endpoint Monitoring Service
An end-to-end POC for scalable and easy onboarding endpoints for Blackbox Exporter monitoring.

This project requires setting up 4 different components: 1) React Frontend, 2) Go Backend, 3) Prometheus, 4) Blackbox Exporter.

## React Frontend
- Dashboard for checking endpoint probe status
- Input Form for sending endpoint records to our backend for storage

### Usage

1. Build Dependencies
```javascript
cd frontend/ && npm install
```

2. Run Frontend
```javascript
npm run start
```

## Go Backend
- Receives endpoints from frontend, saving them in our storage
- Caches the endpoints in regularly intervals
- Exposes the cache through an API to be picked up by Prometheus using HTTP Service Discovery.

### Usage

1. Build Dependencies
```go
go mod tidy
```

2. Run Backend
```go
make run
```

## Prometheus
- Metrics storage
- Runs the prometheus job definition that fetches metrics from blackbox exporter

### Usage
1. Install the binary from [here](https://prometheus.io/download/) or clone the [repo](https://github.com/prometheus/prometheus).

2. Run Prometheus

**Binary**
```shell
./prometheus --config.file=./configs/prometheus.yml
```

**Built from Scratch**
```shell
cd prometheus/ && make build
./prometheus --config.file=./configs/prometheus.yml
```

## Blackbox Exporter
- Performs probes and exports metrics on the statuses of the targets' probes

### Usage
1. Install the binary from [here](https://github.com/prometheus/blackbox_exporter/releases) or clone the [repo](https://github.com/prometheus/blackbox_exporter.git).

2. Run Blackbox Exporter.

**Binary**
```shell
./blackbox_exporter --config.file=./configs/blackbox.yml
```

**Built from Scratch**
```shell
cd blackbox_exporter/ && go run main.go --config.file=./configs/blackbox.yml
```