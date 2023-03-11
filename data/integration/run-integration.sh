#!/usr/bin/env bash

set -x

docker context use default

wget https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.13/grpc_health_probe-linux-amd64
chmod +x ./grpc_health_probe-linux-amd64

docker compose --env-file integration/integration-test.env up --build -d
docker compose --env-file integration/integration-test.env logs --follow &
pid=$!

retcode=1
until [ $retcode -eq 0 ]; do
  ./grpc_health_probe-linux-amd64 -addr localhost:9091
  retcode=$?
  echo "Waiting for server to become healthy..."
  sleep 5
done

echo "Service is healthy, running tests... "

FLIBUSTIER_INTEGRATION=1 FLIBUSERVER_PORT=9091 go test -v integration/flibustier-integration_test.go

echo "Tests finished, cleaning up..."
kill $pid

docker compose --env-file integration/integration-test.env down
docker volume rm flibustier_integration_test_flibudata
rm ./grpc_health_probe-linux-amd64