#!/bin/bash

echo "Doing go test converage"
go test -coverprofile=cov.out

echo "Doing sonar-scanner"
sonar-scanner 