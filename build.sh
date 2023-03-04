#!/bin/bash
export $(grep -v '^#' .env | xargs)
go build -o naming