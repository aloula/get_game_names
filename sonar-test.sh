#!/bin/bash

echo "Doing go test converage"
go test ./test -v -coverpkg=./... -coverprofile=cov.out
go tool cover -html=cov.out -o cov.html

echo "Doing sonar-scanner"
sonar-scanner 