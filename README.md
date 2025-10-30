# Pastebin

A simple pastebin service written in Go.

This project provides a simple API for creating, retrieving, and deleting pastes. It's built with Go, Gin, and GORM, and uses a SQLite database for storage.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

*   Go 1.20 or higher
*   A C compiler (for SQLite)

### Installing

1.  Clone the repository:
    ```bash
    git clone https://github.com/your-username/paste-bin.git
    ```
2.  Navigate to the project directory:
    ```bash
    cd paste-bin
    ```
3.  Install the dependencies:
    ```bash
    go mod tidy
    ```

### Running the application

To run the application, use the following command:

```bash
go run ./cmd/api/main.go
```

The application will start on port `8080` by default.

## API Endpoints

The following API endpoints are available:

*   `POST /pastes`: Create a new paste.
*   `GET /pastes`: Get a list of all pastes.
*   `GET /pastes/:pasteId`: Get a specific paste by its ID.
*   `DELETE /pastes/:pasteId`: Delete a specific paste by its ID.

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
