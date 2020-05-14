#!/bin/bash
ARM_VERSION=6
ARCH_RPI=arm
ARCH_DESKTOP=amd64
PROGRAM_NAME='getgamenames'
 
# Executable name is assumed to be same as current directory name
#EXECUTABLE="${PROGRAM_NAME}_${ARCH}"

#echo "Executable name: ${EXECUTABLE}"
 
echo "Building for Linux ARM (RPI)..."
env GOOS=linux GOARCH=$ARCH_RPI GOARM=$ARM_VERSION go build -o ./builds/getgamenames_rpi getgamenames.go

echo "Building for Linux AMD64..."
env GOOS=linux GOARCH=$ARCH_DESKTOP go build -o ./builds/getgamenames_linux_amd64 getgamenames.go

echo "Building for Windows AMD64..."
env GOOS=windows GOARCH=$ARCH_DESKTOP go build -o ./builds/getgamenames_win_amd64 getgamenames.go
 
