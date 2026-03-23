# Sung-Go

A CLI tool to generate Hexagonal Architecture boilerplate for Go projects.

## Overview

Sung-Go creates a standardized project structure following Hexagonal Architecture (also known as Ports and Adapters) principles. This architecture promotes separation of concerns, testability, and maintainability by isolating business logic from external dependencies.

## Project Structure

The tool generates the following directory structure:

```
project-name/
├── cmd/
│   └── app/              # Application entry point
├── internal/
│   ├── domain/
│   │   ├── entity/       # Business entities and value objects
│   │   └── repository/   # Repository interfaces (ports)
│   ├── application/
│   │   └── service/      # Application services (use cases)
│   ├── ports/
│   │   ├── inbound/      # Driving ports (primary ports)
│   │   └── outbound/     # Driven ports (secondary ports)
│   ├── adapters/
│   │   ├── inbound/      # Driving adapters (e.g., HTTP controllers)
│   │   │   └── http/
│   │   └── outbound/     # Driven adapters (e.g., database, external APIs)
│   │       └── persistence/
│   └── infrastructure/
│       └── database/     # Database configuration and connection
├── configs/              # Configuration files
└── migrations/           # Database migration files
```

## Architecture Layers

### Domain Layer (`internal/domain`)

- **Entity**: Core business objects with business logic
- **Repository**: Interfaces that define data operations (ports)

### Application Layer (`internal/application/service`)

- Application services that orchestrate domain objects
- Implements use cases and business workflows

### Ports Layer (`internal/ports`)

- **Inbound**: Interfaces that the application exposes (driving ports)
- **Outbound**: Interfaces that external systems implement (driven ports)

### Adapters Layer (`internal/adapters`)

- **Inbound**: Implementations of inbound ports (e.g., HTTP handlers, gRPC servers)
- **Outbound**: Implementations of outbound ports (e.g., database repositories, external API clients)

### Infrastructure Layer (`internal/infrastructure`)

- Technical concerns like database connections, configuration loading

## Installation

```bash
# Clone the repository
git clone <repository-url>
cd sung-go

# Build the binary
go build -o sung-go

# Or install directly
go install
```

## Usage

Run the CLI tool:

```bash
./sung-go
```

You will be prompted to enter:

1. **Project name**: The name of your new project
2. **Target directory**: The parent directory where the project will be created

### Example

```
$ ./sung-go
Enter project name: myapi
Enter target directory: ~/projects
Project created at: ~/projects/myapi
```

## Requirements

- Go 1.25.3 or later

<!-- ## Benefits

- **Testability**: Business logic can be tested without external dependencies
- **Maintainability**: Clear separation of concerns
- **Flexibility**: Easy to swap databases, frameworks, or external services
- **Domain-focused**: Core business logic remains isolated and pure

## License

MIT -->
