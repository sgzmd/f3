#!/usr/bin/env bash

set -x

local_run="local_run"

mkdir $local_run
rm "${local_run}/*.sql"
go build -tags fts5 -o "${local_run}/flibustier_server" flibuserver/server/*.go
go build -o "${local_run}/downloader" cmd/downloader/*.go

cp ./mysql2sqlite "${local_run}/"
cp *.sql "${local_run}/"
./create_sqlite_file.sh -f "flibusta.db" -d "${local_run}"
cd $local_run
./flibustier_server --flibusta_db=./flibusta.db --datastore=./datastore.badger --port=9000