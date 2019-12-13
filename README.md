# Auth

## Overview
Auth endpoint 

### Running the service
```shell
make serve
```

### Running Unit Tests
```shell
make test
```

### Linting
```shell
make lint
```

### Docker build
```shell
make docker-build
```

### User Testing
```shell
curl 0.0.0.0:8080/api/v1/health/status
```

### Swagger UI
Swagger ui will be exposed at:
```
0.0.0.0:8080/api/v1/swagger/docs/swaggerui/
```

### Auth endpoint
The auth endpoint checks the header and returns 200 if the headers are correct
```
curl -v -H "eth-id: 0xb4FC6774a95A86134768d81A8eE19Cfbf5A171F6" -H "bam-key: key12345" 0.0.0.0:8080/api/v1/auth/auth
```