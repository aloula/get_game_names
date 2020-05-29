#!/bin/bash

echo "Doing go test converage"
go test ./test -v -coverpkg=./... -coverprofile=cov.out

echo "Doing sonar-scanner"
sonar-scanner 


# sonar-scanner \
#   -Dsonar.projectKey=getgamenames \
#   -Dsonar.sources=. \
#   -Dsonar.host.url=http://localhost:9000 \
#   -Dsonar.login=099d9c61e5f5f8e4abdaa50b86da0911f4b6c02c