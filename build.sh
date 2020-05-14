#!/bin/bash
ARM_VERSION=6
ARCH_RPI=arm
ARCH_DESKTOP=amd64
PROGRAM_NAME_RPI='getgamenames_rpi'
PROGRAM_NAME_LINUX='getgamenames_linux_x64'
PROGRAM_NAME_WINDOWS='getgamenames_win_x64'
 
echo "Building for Linux ARM (RPI)..."
env GOOS=linux GOARCH=$ARCH_RPI GOARM=$ARM_VERSION go build -o ./builds/$PROGRAM_NAME_RPI getgamenames.go

echo "Building for Linux x64..."
env GOOS=linux GOARCH=$ARCH_DESKTOP go build -o ./builds/$PROGRAM_NAME_LINUX getgamenames.go

echo "Building for Windows x64..."
env GOOS=windows GOARCH=$ARCH_DESKTOP go build -o ./builds/$PROGRAM_NAME_WINDOWS getgamenames.go

exit 0 
