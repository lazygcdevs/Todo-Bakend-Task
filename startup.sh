#!/bin/bash

# Azure App Service startup script for Go backend API
# This script is executed when the app service starts

echo "Starting Todo API backend service..."

# Set the working directory
cd /home/site/wwwroot

# Check if the binary exists
if [ -f "main" ]; then
    echo "Found main binary, starting the application..."
    chmod +x main
    ./main
else
    echo "No binary found. Please ensure the application is built before deployment."
    exit 1
fi 