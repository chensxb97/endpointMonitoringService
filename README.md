## Endpoint Monitoring Service
A scalable and easy-to-use self-service portal for managing endpoints for Blackbox Exporter monitoring.

### Frontend
- Dashboard for checking endpoint probe status
- Input Form for sending endpoint records to our backend for storage

### Backend
- Receives endpoints from frontend, saving them in our storage
- Caches the endpoints in regularly intervals
- Exposes the cache through an API to be picked up by Prometheus using HTTP Service Discovery.