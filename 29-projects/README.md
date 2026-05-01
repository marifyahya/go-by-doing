# Category 29: Projects

Build complete applications.

**Total Points: 400**

---

## Project 1: Todo CLI App (100 points)

Build a command-line todo application.

**Features:**
- Add todo
- List todos
- Mark complete
- Delete todo
- Persist to file (JSON)

**Requirements:**
- Use structs for todo
- File-based storage
- Interactive menu

**Structure:**
```
todo/
├── main.go
├── todo.go
├── store.go
└── main.go
```

---

## Project 2: REST API Server (150 points)

Build a REST API with all basics.

**Features:**
- HTTP server
- CRUD endpoints
- JSON requests/responses
- Middleware (logging, CORS)
- In-memory storage

**Endpoints:**
- GET /todos
- POST /todos
- GET /todos/{id}
- PUT /todos/{id}
- DELETE /todos/{id}

**Structure:**
```
api/
├── main.go
├── handler/
├── model/
├── store/
└── middleware/
```

---

## Project 3: Full Stack App (150 points)

Build complete REST API with database.

**Features:**
- PostgreSQL database
- User authentication (JWT)
- Password hashing
- Protected routes
- Redis caching
- Email notifications
- Structured logging
- Background job queue

**Endpoints:**
- POST /auth/register
- POST /auth/login
- GET /todos (protected)
- POST /todos (protected)
- PUT /todos/{id} (protected)
- DELETE /todos/{id} (protected)

**Structure:**
```
full-app/
├── main.go
├── config/
├── handler/
├── model/
├── db/
├── service/
├── middleware/
├── queue/
└── .env
```

**Database Schema:**
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    role VARCHAR(20) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    title VARCHAR(255),
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);
```

**Environment:**
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=learn_go

REDIS_HOST=localhost
REDIS_PORT=6379

SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email
SMTP_PASS=your-password

JWT_SECRET=your-secret-key
```

---

## Bonus Challenges

1. **Rate Limiting** - Limit requests per IP
2. **WebSocket** - Real-time updates
3. **Testing** - Unit tests for handlers
4. **Docker** - Containerize the application
5. **CI/CD** - Setup GitHub Actions