#!/bin/bash
ARM_VERSION=6
ARCH_RPI=arm
ARCH_X86=amd64
ARCH_ARM=arm64
SOURCE_FILE='getgamenames.go'
PROGRAM_NAME='get_game_names'
PROGRAM_NAME_WINDOWS='get_game_names.exe'


echo "Building for Linux x64..."
env GOOS=linux GOARCH=$ARCH_X86 go build -o ./builds/linux_x64/$PROGRAM_NAME $SOURCE_FILE

echo "Building for Linux ARM (RPI)..."
env GOOS=linux GOARCH=$ARCH_RPI GOARM=$ARM_VERSION go build -o ./builds/rpi/$PROGRAM_NAME $SOURCE_FILE

echo "Building for Windows x64..."
env GOOS=windows GOARCH=$ARCH_X86 go build -o ./builds/win_x64/$PROGRAM_NAME_WINDOWS $SOURCE_FILE

echo "Building for MacOS x64..."
env GOOS=darwin GOARCH=$ARCH_X86 go build -o ./builds/mac_x64/$PROGRAM_NAME $SOURCE_FILE

echo "Building for MacOS ARM..."
env GOOS=darwin GOARCH=$ARCH_ARM go build -o ./builds/mac_arm/$PROGRAM_NAME $SOURCE_FILE

exit 0 