# Hello World Webhook in Go

This is a simple "Hello World" application that demonstrates how to create a webhook using Go. The application listens on port 8080 and responds with "Hola Mundo" when accessed through the `/webhook` endpoint.

## Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.

## Getting Started

Follow these steps to run the application:

1. **Clone the Repository:**

    ```sh
    git clone https://github.com/yourusername/hello-world-webhook-go.git
    cd hello-world-webhook-go
    ```

2. **Create the main.go file:**

    Create a file named `main.go` and copy the following code into it:

    ```go
    package main

    import (
        "fmt"
        "net/http"
    )

    func helloWorld(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hola Mundo")
    }

    func main() {
        http.HandleFunc("/webhook", helloWorld)
        http.ListenAndServe(":8080", nil)
    }
    ```

3. **Run the Application:**

    In your terminal, navigate to the directory containing `main.go` and run the following command:

    ```sh
    go run main.go
    ```

