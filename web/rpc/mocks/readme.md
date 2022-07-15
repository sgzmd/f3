To regenerate mocks, run

```shell
mockgen -source=rpc/client.go -package mocks > rpc/mocks/mock_client.go
```

This requires mockgen installed, get it as:

```shell
go install github.com/golang/mock/mockgen@v1.6.0
```