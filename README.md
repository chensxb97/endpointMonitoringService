## Endpoint Monitoring Service
An end-to-end POC for scalable and easy onboarding endpoints for Blackbox Exporter monitoring.

This project requires setting up 4 different components: 1) React Frontend, 2) Go Backend, 3) Blackbox Exporter, 4) Prometheus.

## React Frontend
- Dashboard for checking endpoint probe status
- Input Form for sending endpoint records to our backend for storage

### Usage

1. Build Dependencies
```bash
cd frontend/ && npm install
```

2. Run Frontend
```bash
npm run start
```

3. You should be able to view the following pages: 1) Endpoint Status Dashboard, 2) Endpoint Onboarding Form.

#### Endpoint Status Dashboard

![Dashboard](/assets/Dashboard.png)

#### Endpoint Onboarding Form

![Endpoint Onboarding Form](/assets/EndpointForm.png)


## Go Backend
- Receives endpoints from frontend, saving them in storage. For this example, I am keeping them in memory for simplicity.
- Caches the endpoints in regularly intervals.
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

## Prometheus
- Metric storage
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

3. Validate targets

- Ensure that the endpoints, labels that we have created via the UI are visible on the targets page: `http://localhost:9090/targets`.
![Prometheus Targets](/assets/PrometheusTargets.png)

- Query your metrics via the Prometheus UI.
```bash
# Should return 1 for successful probe, 0 otherwise.
probe_success{instance='http://example.com'}
```
