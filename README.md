# Overview
This project is a command-line journaling application built in Go. 
The primary goal of this software is to demonstrate core Go concepts such as variables, expressions, conditionals, loops, functions, structs, pointers, structuring projects,
handling user input via CLI, and applying the language's modular design through packages. 
The project showcases how to build a simple, interactive journaling tool that can be extended to persist entries and stores entries in JSON files and to other interfaces like web or mobile apps in the future.

The journaling application allows users to create and list journal entries with the capacity to extend to list, update, and delete and view journal entries. The entries are saved in a JSON file, ensuring data persistence across program executions. This project has been designed to be extensible, with a modular architecture that separates the concerns of command-line interaction and data storage.

The purpose of building this software is to improve my understanding of Go’s language syntax, best practices for project organization, and how to implement interactivity in a command-line application. 
Additionally, the project explores designing a future-proof architecture that could be extended to other interfaces, such as web apps.

{Provide a link to your YouTube demonstration. It should be a 4-5 minute demo of the software running and a walkthrough of the code. Focus should be on sharing what you learned about the language syntax.}

[Software Demo Video](https://youtu.be/8qFOscE_rW4)

# Development Environment

## Tools
- Goland
- Go 1.23+
- Git for version control
- Command-line terminal for running the application

## Programming Language and Libraries
- Go (Golang)
- Standard Go libraries (os, fmt, time)
- Google UUID (github.com/google/uuid)
- Custom packages for modular functionality (journal, storage, utils)

# Useful Websites

- [Go Documentation](https://golang.org/doc/)
- [A Tour of Go](https://go.dev/tour/list)
- [Go Modules](https://blog.golang.org/using-go-modules)
- [Golang Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://golang.org/doc/effective_go.html)

# Future Work

- Extend to list, update, and delete and view journal entries
- Persistent data storage using JSON files 
- Integrate a database
- Add search functionality
- Add user authentication 
- Implement web interface