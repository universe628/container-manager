# Container Manager

A containerized file management system with user authentication and file storage capabilities.

## Features

- User Authentication (JWT-based)
- File Upload/Download Management
- Role-based Access Control
- Secure File Storage
- PostgreSQL Database Integration
- Structured Logging (using logrus)

## Project Structure

```
container-manager/
├── cmd/                    # Application entry points
├── database/              # Database migrations and queries
├── internal/              # Internal packages
│   ├── errors/           # Custom error definitions
│   ├── handler/          # HTTP handlers and DTOs
│   ├── infra/           # Infrastructure components
│   ├── middleware/       # HTTP middleware
│   ├── repository/       # Data access layer
│   ├── router/           # HTTP routing
│   ├── schema/           # Domain models
│   ├── service/          # Business logic
│   └── utils/            # Utility functions
└── uploads/              # File storage directory
```

## Prerequisites

- Go 1.20 or later
- PostgreSQL
- Docker (optional)

## Getting Started

1. Clone the repository
```sh
git clone https://github.com/yourusername/container-manager.git
cd container-manager
```

2. Set up the database
```sh
# Run database migrations
migrate -path database/migrations -database "postgresql://your-connection-string" up
```

3. Configuration
Create a configuration file or set environment variables for:
- Database connection
- JWT secret
- File storage path
- Server port

4. Run the application
```sh
go run cmd/server/main.go
```

## API Endpoints

### Authentication
- `POST /auth/signup` - Create new user account
- `POST /auth/login` - Login and receive JWT token

### File Management
- `POST /files/upload` - Upload file
- `GET /files/download/{filename}` - Download file

## Development

### Project Layout
This project follows the standard Go project layout and Clean Architecture principles:

- Domain Layer (`schema/`) - Core business logic and entities
- Service Layer (`service/`) - Use cases and business rules
- Repository Layer (`repository/`) - Data access abstraction
- Infrastructure Layer (`infra/`) - External services and implementations

### Testing
Run tests with:
```sh
go test ./...
```

## License

[Your License]