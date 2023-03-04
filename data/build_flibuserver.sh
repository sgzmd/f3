#!/usr/bin/env bash

go build -o downloader cmd/downloader/main.go 
go build -o flibustier_server flibuserver/server/*.go