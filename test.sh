#!/bin/bash

echo "Doing go test converage"
go test ./test -v -coverpkg=./... -coverprofile=cov.out

echo "Doing sonar-scanner"
sonar-scanner 


# sonar-scanner \
#   -Dsonar.projectKey=getgamenames \
#   -Dsonar.sources=. \
#   -Dsonar.host.url=http://localhost:9000 \
#   -Dsonar.login=8920618f1dc1246322b759676ce31b1152224fd3