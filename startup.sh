#!/bin/bash

# Azure App Service startup script for Go backend API
# This script is executed when the app service starts

echo "Starting Todo API backend service..."

# Set the working directory
cd /home/site/wwwroot

# Build the application
echo "Building Go backend API..."
go build -o main .

# Make the binary executable
chmod +x main

# Start the backend API service
echo "Starting the backend API service on port $PORT..."
./main 