#!/bin/bash
go build -ldflags "-s -w -X 'main.$(grep -v '^#' .env | xargs)'" -o naming -trimpath