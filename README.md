# Overview
This project is a journaling application designed to provide a seamless user experience through both command-line and web-based interfaces including a RESTful API and a frontend interface. 
Users can create, retrieve, update, and delete journal entries, all stored in a persistent SQLite database. 
The web server exposes a RESTful API to facilitate interaction with the journal entries and serves as a way to explore HTTP-based CRUD operations in a real-world scenario.
The frontend application provides a clean and interactive interface for uses to easily view and add journal entries
The journaling application allows users to create and list journal entries with the capacity to extend to list, update, and delete and view journal entries. 
The entries are saved in a local SQLite database, ensuring data persistence across program executions. 
This project has been designed to be extensible, with a modular architecture that separates the concerns of command-line interaction and data storage.

This application was created to deepen my understanding of network server programming in Go, RESTful API development, and database management with SQLite. 
It demonstrates foundational concepts in backend networking, with potential for scaling to more complex systems like fully fledged web application .

[Software Demo Video](https://youtu.be/bg_XjtBb1o4)

# Network Communication
The architecture used in this project is Client-Server. 
The journaling server runs as a standalone HTTP server that can be accessed via HTTP requests from any REST client, such as Postman or cURL
  - **Protocol**: HTTP over TCP
  - **Port**: 8080
  - **Message Format**: JSON payloads for creating and updating entries. JSON responses are provided for each endpoint to allow seamless integration with other services or interfaces.

## REST API Endpoints
  - Create an Entry: **POST /entries** - Expects a JSON payload with title and content.
  - List All Entries: **GET /entries** - Retrieves all journal entries
  - Get a Single Entry: **GET /entries/{id}** - Retrieves a specific entry by its unique id.
  - Update an Entry: **PUT /entries/{id}** - Updates the title and/or content of a specific entry.
  - Delete an Entry: **DELETE /entries/{id}** - Removes a specific entry by its id.

## Web Pages
1. **Home Page**: Displays a list of all journal entries. Each entry can be clicked to view more details or updated. New entries can also be created from this page.
2. **Add New Entry Page**: A form where users can enter the title and content for a new journal entry. After submission, the new entry is added to the database and the user is redirected to the home page.
3. **View Entry Page**: Shows the details of a single journal entry. 

[//]: # (4. From this page, users can edit the entry or delete it.)

## Example requests
```shell
curl http://localhost:8080/entries
```
```shell
curl -X POST -d '{"title":"My Fi
rst Entry","content":"This is my entry"}' \     
-H "Content-Type: application/json" http://localhost:8080/entries
```
```shell
curl -X DELETE http://localhost:8080/entries/entryID
```
```shell
 curl -X PUT  -d '{"title":"second updated title", "content":"second Updated content"}' -H  "Content-Type: application/json" http://localhost:8080/entries/entryID
```
```shell
curl http://localhost:8080/entries/entryID
```


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
- Standard Go libraries (os, fmt, time, database/sql - Go’s standard library package to interact with SQL databases, log, net/http)
- Google UUID ((https://github.com/google/uuid)
- SQLite driver for Go ((https://github.com/mattn/go-sqlite3)
- Gorilla Mux a powerful HTTP router and URL matcher for building Go web servers (https://github.com/gorilla/mux)
- Custom packages for modular functionality (journal, storage, utils)
- Go templates (https://pkg.go.dev/text/template@go1.23.3, https://pkg.go.dev/html/template@go1.23.3)

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
- [Gorilla Mux](https://github.com/gorilla/mux)
- [TailwindCSS Documentation](https://tailwindcss.com/docs)

# Future Work
- Implement user authentication and authorization for secure access.
- Add more sophisticated error handling and user feedback messages.
- Migrate to using PostgreSQL for a more scalable and robust database solution.
- Enhance the styling and layout with more complex TailwindCSS components.
- Explore additional features like filtering, search functionality, and pagination for the entry list.
