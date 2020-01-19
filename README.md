# Auth

## Overview
Service to allow decertralized authorization using smart contracts and blockchain listeners

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

### Api-key endpoint

To easily generate and send an encrypted message use this [message-signer-util](https://github.com/blockchain-abstraction-middleware/sign-message)

```sh
curl -X POST -d '{"hexEthAddress":"0x96216849c49358B10257cb55b28eA603c874b05E",  "hexHash":"0x61769b78c27a41b019fbd77bcf081b95dd399f3403de48e046a2e8c1a6554422","hexSignature":"0xbb1406de26f19b664f3263d19016aadaf797c6c773889cd4f1e1e92d7843008b2436dbffea0a4db6e3d047c86f1615d94f0458affcc9d9e0067d6acba807b0b101"}' 0.0.0.0:8080/api/v1/auth/serve-api-key
```