# Go Backend - Contact Management API

This is a Go backend implementation that replicates the exact functionality of the Node.js Express backend for the Contact Management application.

## Features

- **Clean Architecture**: Organized with proper separation of concerns
- **User Management**: Registration, login, profile management, logout
- **Contact Management**: CRUD operations with search and pagination
- **Address Management**: CRUD operations for contact addresses
- **Authentication**: Token-based authentication middleware
- **Validation**: Request validation using go-playground/validator
- **Configuration**: Viper for configuration management
- **Logging**: Structured logging with Logrus
- **CORS**: Cross-origin resource sharing support
- **Database**: MySQL with raw SQL queries (no ORM)

## Libraries Used

- **gorilla/mux**: HTTP router and URL matcher
- **go-playground/validator**: Struct validation
- **spf13/viper**: Configuration management
- **sirupsen/logrus**: Structured logging
- **gorilla/handlers**: CORS middleware
- **go-sql-driver/mysql**: MySQL driver
- **golang.org/x/crypto**: Password hashing
- **google/uuid**: UUID generation

## Project Structure

```
go-backend/
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   └── config.yaml            # Configuration file
├── internal/
│   ├── config/
│   │   └── config.go          # Configuration management
│   ├── database/
│   │   └── database.go        # Database connection
│   ├── handler/
│   │   ├── user_handler.go    # User HTTP handlers
│   │   ├── contact_handler.go # Contact HTTP handlers
│   │   ├── address_handler.go # Address HTTP handlers
│   │   └── health_handler.go  # Health check handler
│   ├── logger/
│   │   └── logger.go          # Logging configuration
│   ├── middleware/
│   │   ├── auth_middleware.go # Authentication middleware
│   │   └── cors_middleware.go # CORS middleware
│   ├── models/
│   │   ├── user.go           # User models and DTOs
│   │   ├── contact.go        # Contact models and DTOs
│   │   ├── address.go        # Address models and DTOs
│   │   └── response.go       # Response models
│   ├── repository/
│   │   ├── user_repository.go    # User data access
│   │   ├── contact_repository.go # Contact data access
│   │   └── address_repository.go # Address data access
│   ├── router/
│   │   └── router.go         # Route definitions
│   ├── service/
│   │   ├── user_service.go    # User business logic
│   │   ├── contact_service.go # Contact business logic
│   │   └── address_service.go # Address business logic
│   └── utils/
│       ├── validator.go       # Validation utilities
│       └── password.go        # Password utilities
├── Dockerfile                 # Docker configuration
├── docker-compose.yml         # Docker Compose configuration
├── init.sql                   # Database initialization
├── go.mod                     # Go module dependencies
└── README.md                  # This file
```

## API Endpoints

### Public Endpoints
- `POST /api/users` - Register new user
- `POST /api/users/login` - User login
- `GET /ping` - Health check

### Protected Endpoints (require Authorization header)

#### User Management
- `GET /api/users/current` - Get current user
- `PATCH /api/users/current` - Update current user
- `DELETE /api/users/logout` - Logout user

#### Contact Management
- `POST /api/contacts` - Create contact
- `GET /api/contacts/{id}` - Get contact by ID
- `PUT /api/contacts/{id}` - Update contact
- `DELETE /api/contacts/{id}` - Delete contact
- `GET /api/contacts` - Search contacts (with pagination)

#### Address Management
- `POST /api/contacts/{contactId}/addresses` - Create address
- `GET /api/contacts/{contactId}/addresses/{addressId}` - Get address
- `PUT /api/contacts/{contactId}/addresses/{addressId}` - Update address
- `DELETE /api/contacts/{contactId}/addresses/{addressId}` - Delete address
- `GET /api/contacts/{contactId}/addresses` - List addresses

## Configuration

The application uses `config/config.yaml` for configuration:

```yaml
server:
  port: 3000
  host: localhost

database:
  host: localhost
  port: 3306
  username: root
  password: root
  name: belajar_vuejs_contact_management

logging:
  level: info
  format: json
```

## Running the Application

### Prerequisites
- Go 1.21 or higher
- MySQL 8.0 or higher

### Local Development

1. **Clone and navigate to the project:**
   ```bash
   cd go-backend
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Set up the database:**
   - Create MySQL database: `belajar_vuejs_contact_management`
   - Run the SQL scripts from `init.sql`

4. **Update configuration:**
   - Modify `config/config.yaml` with your database credentials

5. **Run the application:**
   ```bash
   go run cmd/main.go
   ```

### Using Docker

1. **Build and run with Docker Compose:**
   ```bash
   docker-compose up --build
   ```

This will start both the Go backend and MySQL database.

## Testing

The API endpoints are compatible with the existing frontend and can be tested using the same HTTP requests as the Node.js version.

Example test requests:

```bash
# Register user
curl -X POST http://localhost:3000/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"password","name":"Test User"}'

# Login
curl -X POST http://localhost:3000/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"password"}'

# Get current user (replace TOKEN with actual token)
curl -X GET http://localhost:3000/api/users/current \
  -H "Authorization: TOKEN"
```

## Key Features

1. **Exact API Compatibility**: All endpoints return the same response format as the Node.js version
2. **Clean Architecture**: Proper separation of concerns with layers
3. **No ORM**: Uses raw SQL queries for better performance and control
4. **Comprehensive Validation**: Request validation with detailed error messages
5. **Structured Logging**: JSON-formatted logs for better observability
6. **Docker Support**: Easy deployment with Docker and Docker Compose
7. **Configuration Management**: Flexible configuration with Viper
8. **Security**: Password hashing and token-based authentication

This Go backend provides the exact same functionality as the Node.js version while following Go best practices and conventions.