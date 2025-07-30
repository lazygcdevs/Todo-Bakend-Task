#!/bin/bash

# Azure App Service startup script for Go backend API
# This script is executed when the app service starts

echo "Starting Todo API backend service..."

# Set the working directory
cd /home/site/wwwroot

# Check if the binary exists
if [ -f "main.exe" ]; then
    echo "Found main.exe, starting the application..."
    chmod +x main.exe
    ./main.exe
elif [ -f "main" ]; then
    echo "Found main binary, starting the application..."
    chmod +x main
    ./main
else
    echo "Building the application..."
    go build -o main .
    chmod +x main
    ./main
fi 