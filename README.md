# Overview
This project is a command-line journaling application built in Go that integrates with a SQLite relational database to store journal entries. 
The primary goal of this software is to deepen my understanding of Go (concepts such as variables, expressions, conditionals, loops, functions, structs, pointers, structuring projects,
handling user input via CLI, and applying the language's modular design through packages) and persistent data storage through its integration with relation databases like SQL. 
The project showcases how to build a simple, interactive journaling tool that can be extended to persist entries and stores entries in JSON files and to other interfaces like web or mobile apps in the future.

The journaling application allows users to create and list journal entries with the capacity to extend to list, update, and delete and view journal entries. The entries are saved in a local SQLite database, ensuring data persistence across program executions. This project has been designed to be extensible, with a modular architecture that separates the concerns of command-line interaction and data storage.

This journaling application can be easily extended for web applications or other interfaces in the future due to its modular architecture.{Provide a link to your YouTube demonstration. It should be a 4-5 minute demo of the software running and a walkthrough of the code. Focus should be on sharing what you learned about the language syntax.}

[Software Demo Video](https://youtu.be/K_7iwzL7ORE)

# Relational Database
For this project, I used SQLite, a lightweight SQL relational database that is ideal for local development or small-scale applications. The database stores all journal entries, each with a unique ID, title, content, and timestamps (created and updated).

## Database Structure:
- Table: journal_entries
    - **id** (TEXT) - Primary Key, a unique identifier for each entry.
    - **title** (TEXT) - The title of the journal entry.
    - **content** (TEXT) - The content or body of the journal entry.
    - **created** (TIMESTAMP) - The timestamp of when the entry was created.
    - **updated** (TIMESTAMP) - The timestamp of when the entry was last updated.

Whenever the application starts, it checks if the journal_entries table exists and creates it if not, ensuring seamless operation even on first use.

# Development Environment

## Tools
- Goland
- Go 1.23+
- Git for version control
- Command-line terminal for running the application

## Programming Language and Libraries
- Go (Golang)
- Standard Go libraries (os, fmt, time, database/sql - Go’s standard library package to interact with SQL databases)
- Google UUID (github.com/google/uuid)
- SQLite driver for Go (github.com/mattn/go-sqlite3)
- Custom packages for modular functionality (journal, storage, utils)

# Useful Websites

- [Go Documentation](https://golang.org/doc/)
- [A Tour of Go](https://go.dev/tour/list)
- [Go Modules](https://blog.golang.org/using-go-modules)
- [Golang Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [SQLite Documentation](https://www.sqlite.org/docs.html)
- [Tutorial: Accessing a relational database](https://go.dev/doc/tutorial/database-access)
- [Go SQLite3 Driver Documentation](https://pkg.go.dev/github.com/mattn/go-sqlite3)
- [Using SQLite from Go](https://practicalgobook.net/posts/go-sqlite-no-cgo/)
- [Golang SQLite `database/sql`](https://earthly.dev/blog/golang-sqlite/)

# Future Work

- Add search functionality
- Add user authentication 
- Implement web interface