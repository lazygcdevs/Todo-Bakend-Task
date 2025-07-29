#!/bin/bash

echo "🚀 Todo API Setup Script"
echo "========================"

# Check if .env file exists
if [ ! -f ".env" ]; then
    echo "📝 Creating .env file with Azure Cosmos DB configuration..."
    cat > .env << EOF
MONGODB_URI=mongodb+srv://nipuna:IyeFY1jdK.S1ks@dbtodo.mongocluster.cosmos.azure.com/?tls=true&authMechanism=SCRAM-SHA-256&retrywrites=false&maxIdleTimeMS=120000
DATABASE_NAME=todoapp
COLLECTION_NAME=todos
PORT=8080
COOKIE_NAME=todo_user_id
EOF
    echo "✅ .env file created successfully!"
else
    echo "✅ .env file already exists"
fi

# Install dependencies
echo "📦 Installing Go dependencies..."
go mod tidy

# Build the application
echo "🔨 Building the application..."
go build -o todo-api

echo "🎉 Setup complete!"
echo ""
echo "To start the API server, run:"
echo "  ./todo-api"
echo "or"
echo "  go run main.go"
echo ""
echo "The API will be available at: http://localhost:8080"
echo "Health check: http://localhost:8080/health" 