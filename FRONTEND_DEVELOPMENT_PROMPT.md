# Frontend Development Prompt: Todo List Application

## Project Overview
Create a modern, responsive todo list web application that connects to an existing REST API. The application should provide a clean, intuitive interface for managing personal todo tasks with automatic user identification (no login required).

## API Integration Details

### Base Configuration
- **API Base URL:** `http://localhost:8080`
- **Authentication:** Cookie-based (automatic, no login required)
- **CORS:** Enabled for all origins with credentials support

### Critical Implementation Requirements
1. **Always include credentials:** Set `credentials: 'include'` in all fetch requests or `withCredentials: true` for axios
2. **Content-Type:** Use `application/json` for POST/PUT requests
3. **Error handling:** Implement proper error handling for all API calls
4. **Loading states:** Show loading indicators during API operations

### Available API Endpoints

#### 1. Health Check
```
GET /health
Response: {"status": "ok", "message": "Todo API is running"}
```

#### 2. Get All Todos
```
GET /api/v1/todos
Response: {"todos": [Todo...]}
```

#### 3. Create Todo
```
POST /api/v1/todos
Body: {"title": "string (required)", "description": "string (optional)"}
Response: {"todo": Todo}
```

#### 4. Update Todo
```
PUT /api/v1/todos/{id}
Body: {"title": "string (optional)", "description": "string (optional)", "completed": "boolean (optional)"}
Response: {"todo": Todo}
```

#### 5. Delete Todo
```
DELETE /api/v1/todos/{id}
Response: {"message": "Todo deleted successfully"}
```

### Data Model
```typescript
interface Todo {
  id: string;                // MongoDB ObjectID
  user_id: string;          // UUID (auto-generated)
  title: string;            // Todo title
  description: string;      // Todo description
  completed: boolean;       // Completion status
  created_at: string;       // ISO 8601 datetime
  updated_at: string;       // ISO 8601 datetime
}
```

### API Client Example (Use as Reference)
```javascript
const API_BASE = 'http://localhost:8080/api/v1';

async function apiCall(endpoint, options = {}) {
  const response = await fetch(`${API_BASE}${endpoint}`, {
    credentials: 'include', // CRITICAL: Always include cookies
    headers: {
      'Content-Type': 'application/json',
      ...options.headers
    },
    ...options
  });
  
  if (!response.ok) {
    throw new Error(`API Error: ${response.status}`);
  }
  
  return response.json();
}

// Usage examples:
const todos = await apiCall('/todos');
const newTodo = await apiCall('/todos', {
  method: 'POST',
  body: JSON.stringify({ title: 'New Todo', description: 'Description' })
});
```

## UI/UX Requirements

### Design Guidelines
- **Modern and clean:** Use contemporary design patterns
- **Responsive:** Work well on desktop, tablet, and mobile
- **Intuitive:** Clear visual hierarchy and easy navigation
- **Accessible:** Proper ARIA labels, keyboard navigation
- **Fast:** Optimistic updates where possible

### Required Features

#### 1. Todo List View
- Display all todos in a clean list format
- Show title, description (if any), and completion status
- Include created/updated timestamps
- Empty state message when no todos exist
- Loading state while fetching data

#### 2. Add Todo Functionality
- Input field for title (required)
- Textarea for description (optional)
- Add button (disabled when title is empty)
- Clear form after successful creation
- Show success/error messages

#### 3. Todo Item Actions
- **Toggle completion:** Click/tap to mark complete/incomplete
- **Edit inline:** Double-click to edit title/description
- **Delete:** Delete button with confirmation
- **Visual distinction:** Completed todos should look different (strikethrough, dimmed, etc.)

#### 4. Edit Todo
- Inline editing or modal for updating title/description
- Save/cancel buttons
- Validate required fields
- Update UI immediately on success

#### 5. Status Management
- Loading indicators for all async operations
- Error messages for failed operations
- Success confirmations for completed actions
- Retry functionality for failed requests

### Suggested UI Layout

#### Header Section
- Application title "Todo List"
- Optional subtitle or user indicator
- Add new todo form (title input + description textarea + add button)

#### Main Content
- List of todo items
- Each item should display:
  - Checkbox or toggle for completion
  - Todo title (with strikethrough if completed)
  - Description (if present)
  - Edit and delete action buttons
  - Timestamp information
- Empty state when no todos

#### Footer (Optional)
- Total todos count
- Completed/remaining count
- Clear completed button (if any completed todos exist)

## Technical Requirements

### Framework Preference
Choose one of:
- **React** (with hooks)
- **Vue.js** (composition API)
- **Vanilla JavaScript** (modern ES6+)
- **Angular** (latest version)

### State Management
- Use appropriate state management for the chosen framework
- Manage loading states, error states, and todo data
- Consider local state vs. component props appropriately

### Styling Approach
Choose one of:
- **CSS Modules**
- **Styled Components** (React)
- **Tailwind CSS**
- **Regular CSS/SCSS**
- **Material-UI/Chakra-UI** (component library)

### Key Implementation Details

#### Error Handling
```javascript
// Example error handling pattern
try {
  setLoading(true);
  const result = await createTodo(title, description);
  setTodos(prev => [...prev, result]);
  setSuccessMessage('Todo created successfully');
  clearForm();
} catch (error) {
  setErrorMessage('Failed to create todo. Please try again.');
} finally {
  setLoading(false);
}
```

#### Loading States
- Show spinners or skeletons during API calls
- Disable form elements during submission
- Use optimistic updates where appropriate

#### Form Validation
- Required field validation for todo title
- Character limits if desired
- Real-time validation feedback

## Implementation Steps

### Phase 1: Basic Setup
1. Set up project with chosen framework
2. Create basic component structure
3. Implement API client with proper credential handling
4. Test API connectivity

### Phase 2: Core Functionality
1. Implement todo list display
2. Add create todo functionality
3. Add toggle completion feature
4. Add delete functionality

### Phase 3: Enhanced Features
1. Implement edit functionality
2. Add error handling and loading states
3. Implement responsive design
4. Add animations/transitions

### Phase 4: Polish
1. Add accessibility features
2. Implement empty states
3. Add success/error messaging
4. Performance optimization

## Testing Instructions

### Manual Testing Checklist
- [ ] App loads without errors
- [ ] Can create new todos
- [ ] Can view all todos
- [ ] Can mark todos as complete/incomplete
- [ ] Can edit existing todos
- [ ] Can delete todos
- [ ] Error handling works for network failures
- [ ] Loading states are visible
- [ ] Responsive design works on mobile
- [ ] Cookies are automatically handled (no manual auth)

### API Testing
Use these curl commands to verify API connectivity:
```bash
# Health check
curl http://localhost:8080/health

# Get todos (will auto-generate user)
curl -c cookies.txt http://localhost:8080/api/v1/todos

# Create todo
curl -b cookies.txt -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Todo","description":"Test description"}'
```

## Real API Testing Examples

Here are actual working examples from API testing:

### 1. Health Check
**Request:**
```bash
curl -s http://localhost:8080/health
```

**Response:**
```json
{"message":"Todo API is running","status":"ok"}
```

### 2. Get Todos (Empty List)
**Request:**
```bash
curl -s -c cookies.txt http://localhost:8080/api/v1/todos
```

**Response:**
```json
{"todos":[]}
```

### 3. Create Todo
**Request:**
```bash
curl -s -b cookies.txt -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Todo","description":"This is a test todo item"}'
```

**Response:**
```json
{
  "todo": {
    "id": "6888208b8a67a0d68ef14079",
    "user_id": "5909f0cd-e749-4050-81b3-04ade6d08e55",
    "title": "Test Todo",
    "description": "This is a test todo item",
    "completed": false,
    "created_at": "2025-07-29T01:14:51.473Z",
    "updated_at": "2025-07-29T01:14:51.473Z"
  }
}
```

### 4. Get Todos (With Data)
**Request:**
```bash
curl -s -b cookies.txt http://localhost:8080/api/v1/todos
```

**Response:**
```json
{
  "todos": [
    {
      "id": "6888208b8a67a0d68ef14079",
      "user_id": "5909f0cd-e749-4050-81b3-04ade6d08e55",
      "title": "Test Todo",
      "description": "This is a test todo item",
      "completed": false,
      "created_at": "2025-07-29T01:14:51.473Z",
      "updated_at": "2025-07-29T01:14:51.473Z"
    }
  ]
}
```

### 5. Update Todo
**Request:**
```bash
curl -s -b cookies.txt -X PUT http://localhost:8080/api/v1/todos/6888208b8a67a0d68ef14079 \
  -H "Content-Type: application/json" \
  -d '{"completed":true,"description":"Updated description"}'
```

**Response:**
```json
{
  "todo": {
    "id": "6888208b8a67a0d68ef14079",
    "user_id": "5909f0cd-e749-4050-81b3-04ade6d08e55",
    "title": "Test Todo",
    "description": "Updated description",
    "completed": true,
    "created_at": "2025-07-29T01:14:51.473Z",
    "updated_at": "2025-07-29T01:15:02.234Z"
  }
}
```

### 6. Delete Todo
**Request:**
```bash
curl -s -b cookies.txt -X DELETE http://localhost:8080/api/v1/todos/6888208b8a67a0d68ef14079
```

**Response:**
```json
{"message":"Todo deleted successfully"}
```

### 7. Verify Deletion
**Request:**
```bash
curl -s -b cookies.txt http://localhost:8080/api/v1/todos
```

**Response:**
```json
{"todos":[]}
```

## Server Logs During Testing

The API server showed these logs during our testing session:
```
[GIN] 2025/07/28 - 21:14:40 | 200 |      69.485µs |             ::1 | GET      "/health"
[GIN] 2025/07/28 - 21:14:48 | 200 |  282.512359ms |             ::1 | GET      "/api/v1/todos"
[GIN] 2025/07/28 - 21:14:51 | 201 |  327.240392ms |             ::1 | POST     "/api/v1/todos"
[GIN] 2025/07/28 - 21:14:55 | 200 |   274.28769ms |             ::1 | GET      "/api/v1/todos"
[GIN] 2025/07/28 - 21:15:02 | 200 |  472.843733ms |             ::1 | PUT      "/api/v1/todos/6888208b8a67a0d68ef14079"
[GIN] 2025/07/28 - 21:15:07 | 200 |  240.483945ms |             ::1 | DELETE   "/api/v1/todos/6888208b8a67a0d68ef14079"
[GIN] 2025/07/28 - 21:15:11 | 200 |  291.491471ms |             ::1 | GET      "/api/v1/todos"
```

**Key Observations:**
- All operations completed successfully (200/201 status codes)
- Response times are reasonable (69µs to 472ms)
- User ID remains consistent across all operations (`5909f0cd-e749-4050-81b3-04ade6d08e55`)
- MongoDB ObjectIDs are properly generated (`6888208b8a67a0d68ef14079`)
- Timestamps are in ISO 8601 format
- Cookie-based authentication works seamlessly

## CORS Configuration & Troubleshooting

### Important CORS Note
The API has been configured to work with `credentials: 'include'` by explicitly allowing specific origins instead of using wildcard (`*`). This is required because browsers block wildcard origins when credentials are included.

**Allowed Origins:**
- `http://localhost:3000` (React default)
- `http://localhost:5173` (Vite default) 
- `http://localhost:8080` (API same-origin)
- `http://127.0.0.1:3000`
- `http://127.0.0.1:5173`
- `http://127.0.0.1:8080`

### Common CORS Issues & Solutions

#### Error: "Response to preflight request doesn't pass access control check"
**Problem:** Using `AllowAllOrigins = true` with `credentials: 'include'`

**Solution:** The API has been fixed to explicitly allow specific origins.

#### If you see CORS errors:
1. **Check your frontend URL** - Make sure it matches one of the allowed origins above
2. **Always include credentials** - Use `credentials: 'include'` in all fetch requests
3. **Check server logs** - Look for OPTIONS requests (preflight) getting 204 responses

### API Server CORS Configuration
```go
config.AllowOrigins = []string{
    "http://localhost:3000",  // React default
    "http://localhost:5173",  // Vite default
    "http://localhost:8080",  // Same origin
    "http://127.0.0.1:3000",
    "http://127.0.0.1:5173", 
    "http://127.0.0.1:8080",
}
config.AllowCredentials = true
config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
```

## Success Criteria
- [ ] All CRUD operations work correctly
- [ ] User can manage todos without any authentication
- [ ] UI is responsive and accessible
- [ ] Error handling provides clear feedback
- [ ] Loading states prevent confusion
- [ ] Application is production-ready

## Additional Notes
- The API automatically handles user identification via cookies
- No login/registration system is needed
- Users get a unique ID automatically on first visit
- Cookie expires after 24 hours but renews on each request
- All todos are automatically filtered by user ID
- Database and collections are created automatically on first use

Please create a complete, production-ready todo application that integrates seamlessly with this API. 