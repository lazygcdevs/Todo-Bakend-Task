# Todo API - Golang with Azure Cosmos DB

A RESTful API for managing todo lists built with Go, Gin framework, and Azure Cosmos DB (MongoDB API). Users are automatically identified through cookies - no login required!

## Features

- 🚀 **Cookie-based Authentication**: Auto-generates user IDs, no login required
- 📱 **CRUD Operations**: Create, Read, Update, Delete todos
- ☁️ **Azure Cosmos DB**: MongoDB API for cloud storage
- 🌐 **CORS Enabled**: Accessible from any frontend origin
- ⚡ **Gin Framework**: Fast and lightweight HTTP router

## Quick Start

### 1. Environment Setup

Create a `.env` file in the project root with the following variables:

```env
MONGODB_URI=mongodb+srv://usenrmae@passworddbtodo.mongocluster.cosmos.azure.com/?tls=true&authMechanism=SCRAM-SHA-256&retrywrites=false&maxIdleTimeMS=120000
DATABASE_NAME=todoapp
COLLECTION_NAME=todos
PORT=8080
COOKIE_NAME=todo_user_id
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the API

```bash
go run main.go
```

The API will start on `http://localhost:8080`

## API Endpoints

### Health Check
- **GET** `/health` - Check if API is running

### Todo Operations
All endpoints automatically handle user identification via cookies.

- **GET** `/api/v1/todos` - Get all todos for the user
- **POST** `/api/v1/todos` - Create a new todo
- **PUT** `/api/v1/todos/:id` - Update a specific todo
- **DELETE** `/api/v1/todos/:id` - Delete a specific todo

## Request/Response Examples

### Create Todo
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Complete the todo API project"
  }'
```

**Response:**
```json
{
  "todo": {
    "id": "507f1f77bcf86cd799439011",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Learn Go",
    "description": "Complete the todo API project",
    "completed": false,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

### Get All Todos
```bash
curl http://localhost:8080/api/v1/todos
```

**Response:**
```json
{
  "todos": [
    {
      "id": "507f1f77bcf86cd799439011",
      "user_id": "550e8400-e29b-41d4-a716-446655440000",
      "title": "Learn Go",
      "description": "Complete the todo API project",
      "completed": false,
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    }
  ]
}
```

### Update Todo
```bash
curl -X PUT http://localhost:8080/api/v1/todos/507f1f77bcf86cd799439011 \
  -H "Content-Type: application/json" \
  -d '{
    "completed": true
  }'
```

### Delete Todo
```bash
curl -X DELETE http://localhost:8080/api/v1/todos/507f1f77bcf86cd799439011
```

## Data Models

### Todo
```go
type Todo struct {
    ID          primitive.ObjectID `json:"id"`
    UserID      string             `json:"user_id"`
    Title       string             `json:"title"`
    Description string             `json:"description"`
    Completed   bool               `json:"completed"`
    CreatedAt   time.Time          `json:"created_at"`
    UpdatedAt   time.Time          `json:"updated_at"`
}
```

## How Authentication Works

1. When a user first makes a request, the API automatically generates a unique UUID
2. This UUID is stored in a secure HTTP cookie with 24-hour expiration
3. All subsequent requests use this cookie to identify the user
4. No login/registration required - users just start using the app!

## Architecture

```
├── main.go              # Entry point and server setup
├── models/
│   └── todo.go         # Data structures
├── handlers/
│   └── todo.go         # API request handlers
├── middleware/
│   └── auth.go         # Cookie-based authentication
├── database/
│   └── connection.go   # Azure Cosmos DB connection
└── go.mod              # Dependencies
```

## Dependencies

- **Gin**: HTTP router and middleware
- **MongoDB Driver**: Azure Cosmos DB connectivity
- **UUID**: User ID generation
- **GoDotEnv**: Environment variable loading
- **CORS**: Cross-origin resource sharing

## Production Considerations

1. **HTTPS**: Enable secure cookies in production
2. **Environment Variables**: Use secure secret management
3. **Error Logging**: Add structured logging
4. **Rate Limiting**: Implement request throttling
5. **Validation**: Add input sanitization
6. **Monitoring**: Add health checks and metrics

## Development

### Run with Hot Reload
```bash
# Install air for hot reload
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

### Testing
```bash
go test ./...
```

## License

MIT License 
