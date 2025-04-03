# Golang Basic: Your First Steps with Go

## Welcome to the Golang Basic Repository!

This repository serves as a fundamental guide for those looking to start or refresh their knowledge of the Go programming language (Golang).

## Installation Guide: Step-by-Step

Before embarking on your coding adventure with Go, ensure that Go is installed on your system. Follow these steps:

1.  **Download Go:**
    * Visit the [official Go website](https://golang.org/dl/) and download the latest version suitable for your operating system.

2.  **Verify Installation:**
    * Open your terminal (on Linux/macOS) or Command Prompt (on Windows).
    * Run the following command to confirm that Go is installed correctly and view its version:
        ```bash
        go version
        ```
        You should see output indicating the installed Go version.

3.  **Running Go Programs:**
    * To execute a Go program file (`.go`), use the command:
        ```bash
        go run program_name.go
        ```
        Replace `program_name.go` with the name of your Go file.

    * **Faster Development with Hot Reload:**
        For a more efficient development experience, especially for web applications, consider using *hot reloading*. A popular tool for this is **Air**. If you haven't installed it, run:
        ```bash
        go install [github.com/cosmtrek/air@latest](https://www.google.com/search?q=https://github.com/cosmtrek/air%40latest)
        ```
        Then, in your project directory, run `air`. This will automatically restart your application when you save changes.

4.  **Compiling Go Programs:**
    * To compile a Go program into an executable file that can be run without the `go run` command, use:
        ```bash
        go build program_name.go
        ```
        This will generate an executable file in the same directory as your program, named after your program (or `program_name.exe` on Windows).

## Official Go Documentation Resources

For a deeper understanding and comprehensive information about Go, don't hesitate to explore the following official resources:

* **Go Documentation:** [https://golang.org/doc/](https://golang.org/doc/) - The complete reference for the Go language, standard library, and tools.
* **Go Tour:** [https://tour.golang.org/](https://tour.golang.org/) - An interactive tutorial that helps you learn Go through hands-on exercises in your browser.
* **Package Documentation:** [https://pkg.go.dev/](https://pkg.go.dev/) - A comprehensive catalog of all standard and third-party Go packages, along with their documentation and usage examples.
