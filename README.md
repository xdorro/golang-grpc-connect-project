# golang-grpc-base-project

## Getting started

Setup GO PRIVATE

Windows:
```
go env -w GOPRIVATE=github.com/*
```

Linux:
```
export GOPRIVATE=github.com/*
```

Get all modules
```
env GIT_TERMINAL_PROMPT=1 go mod tidy
```

## Example

The service is running on http://localhost:8088. To make an RPC with cURL,
using the Connect protocol:

```bash
curl --header "Content-Type: application/json" \
    --data '{
        "email": "admin@gmail.com",
        "password": "123456"
    }' \
    localhost:5000/auth.v1.AuthService/Login
```

To make the same RPC, but using [`grpcurl`][grpcurl] and the gRPC protocol:

```bash
grpcurl \
    -plaintext \
    -d '{
        "email": "admin@gmail.com",
        "password": "123456"
    }' \
    localhost:5000 \
    auth.v1.AuthService/Login
```