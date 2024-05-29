# Final SSE Implementation

This directory contains the final implementation of the SSE server with simulated word generation.

## Requirements

- Golang
- Echo framework

## How to Run

1. Ensure you have Go installed on your system.

2. Clone this repository to your local machine:
    ```sh
    git clone https://github.com/yourusername/streaming-text-generation.git
    ```

3. Navigate to the `final-implementation` directory:
    ```sh
    cd streaming-text-generation/final-implementation
    ```

4. Run `go mod tidy` to ensure all dependencies are up to date:
    ```sh
    go mod tidy
    ```

5. Run the server:
    ```sh
    go run server.go
    ```

6. Open `index.html` in a web browser to see the SSE in action.