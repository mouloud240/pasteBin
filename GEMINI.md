# Project Overview

This project is a simple pastebin service written in Go. It uses the Gin framework for handling HTTP requests and GORM as an ORM for interacting with a SQLite database. The service allows users to create, retrieve, and delete pastes. Pastes can be protected with a password, have an expiration date, and a maximum number of views.

# Building and Running

## Prerequisites

*   Go 1.20 or higher
*   A C compiler (for SQLite)

## Running the application

To run the application, you can use the following command:

```bash
go run ./cmd/api/main.go
```

The application will start on port 8080 by default.

## Building the application

To build the application, you can use the following command:

```bash
go build -o pastebin ./cmd/api/main.go
```

This will create an executable file named `pastebin` in the root directory.

# Development Conventions

## Code Style

The project follows the standard Go code style. It is recommended to use `gofmt` to format your code before committing.

## API Endpoints

The following API endpoints are available:

*   `POST /pastes`: Create a new paste.
*   `GET /pastes`: Get a list of all pastes.
*   `GET /pastes/:pasteId`: Get a specific paste by its ID.
*   `DELETE /pastes/:pasteId`: Delete a specific paste by its ID.

## Database

The project uses a SQLite database to store the pastes. The database file is named `app.db` and is located in the root directory. The database schema is defined in the `internal/database/models` directory.
