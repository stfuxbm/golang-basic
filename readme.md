## Golang Basic Repository

This repository is a basic guide for those of you who want to start learning or refresh your knowledge of the fundamentals of the Go (Golang) programming language.

## How to Install Go

Before you start coding with Go, make sure Go is installed on your computer. Here's how:

1.  **Download Go:**
    * Visit the official Go website at [https://golang.org/dl/](https://golang.org/dl/) and download the latest version that matches your computer's operating system (Windows, macOS, or Linux).

2.  **Check Installation:**
    * Open your terminal (on Linux/macOS) or Command Prompt (on Windows).
    * Type the command below and press Enter to see if Go is installed and its version:
        ```bash
        go version
        ```
        You will see information about the installed Go version. **Currently, based on this repository, the Go version being used is:**
        ```
        go version go1.22.4 darwin/amd64
        ```
        (This indicates Go version 1.22.4 running on macOS (darwin) with an AMD64 architecture.)

3.  **Running Go Programs:**
    * To run a Go program file (which has the `.go` extension), use this command in the terminal:
        ```bash
        go run your_program_name.go
        ```
        Replace `your_program_name.go` with the name of your Go file.

    * **Faster Development with Air (Optional):**
        If you are building web applications and want changes to be reflected immediately without having to rerun continuously, you can use a tool called **Air**. Here's how to install it:
        ```bash
        go install [github.com/cosmtrek/air@latest](https://www.google.com/search?q=https://github.com/cosmtrek/air%40latest)
        ```
        After installation, open the terminal in your project folder and run:
         ```bash
        air
        ```
4.  **Turning Go Programs into Applications:**
    * To turn your Go program into an executable application file that can be run directly without the `go run` command, use this command:
        ```bash
        go build your_program_name.go
        ```
        The application file will appear in the same folder with your program's name (e.g., `your_program_name` on Linux and macOS, or `your_program_name.exe` on Windows).

## Official Go Learning Resources

If you want to learn more deeply and comprehensively about Go, don't hesitate to check out these official resources:

* **Go Documentation:** [https://golang.org/doc/](https://golang.org/doc/) - A complete reference for the Go language.
* **Go Tour:** [https://tour.golang.org/](https://tour.golang.org/) - An interactive tutorial to learn Go directly in your browser.
* **Package Documentation:** [https://pkg.go.dev/](https://pkg.go.dev/) - A catalog of all standard and third-party libraries (packages) in Go.