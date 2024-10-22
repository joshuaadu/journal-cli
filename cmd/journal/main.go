package main

import (
	"bufio"
	"fmt"
	"journal/pkg/journal"
	"journal/pkg/storage"
	"os"
	"strings"
)

func main() {
	// Create a new journal instance
	//journalInstance := journal.NewJournal()

	// SQLite file for storing journal entries
	dbFileName := "journal.db"

	// Initialize SQLite storage
	sqliteStorage, err := storage.NewSQLiteStorage(dbFileName)

	if err != nil {
		fmt.Println("Error initializing SQLite: ", err)
		return
	}

	// Create a new journal instance using SQLite
	journalInstance := journal.NewJournal(sqliteStorage)

	// Check command line arguments
	if len(os.Args) < 2 {
		fmt.Println("usage: journal [command] [arguments]")
		return
	}

	// Determine the command
	command := os.Args[1]

	switch command {
	case "create":
		if len(os.Args) < 4 {
			fmt.Println("usage: journal create [title] [content]")
			return
		}

		title := os.Args[2]
		content := os.Args[3]

		entry := journalInstance.CreateEntry(title, content)
		fmt.Println("Created entry: %s\n", entry.ID)

	case "list":
		// List all journal entries
		entries := journalInstance.ListEntries()
		if len(entries) < 1 {
			fmt.Println("No entries found.")
			return
		}
		for _, entry := range entries {
			fmt.Printf(" ID: %s\n Title: %s\n Content: %s\n Created: %s\n\n", entry.ID, entry.Title, entry.Content, entry.Created)

		}
	case "":
	case "interactive":
		// Create a scanner to read input from the terminal
		scanner := bufio.NewScanner(os.Stdin)
		var title, content string
		fmt.Println("You are in the Journal program in interactive mode")
		for {
			fmt.Println("usage: create | list | get | update | delete")
			fmt.Println("Type 'exit' to quit.")
			// Prompt the user for input
			fmt.Print("> ")

			// Read user input
			if scanner.Scan() {
				input := scanner.Text()
				input = strings.TrimSpace(input)

				// Check for the exit command
				if input == "exit" {
					fmt.Println("Exiting...")
					break
				}

				// handle other commands here
				switch input {
				case "create":
					fmt.Println("Enter entry title")
					scanner.Scan()
					title = scanner.Text()
					fmt.Println("Enter entry content")
					scanner.Scan()
					content = scanner.Text()
					entry := journalInstance.CreateEntry(title, content)

					fmt.Println("Created entry: %s\n", entry.ID)

				case "list":
					entries := journalInstance.ListEntries()
					if len(entries) < 1 {
						fmt.Println("No entries found.")
					}
					for _, entry := range entries {
						fmt.Printf(" ID: %s\n Title: %s\n Content: %s\n Created: %s\n\n", entry.ID, entry.Title, entry.Content, entry.Created)
					}

				case "get":
					fmt.Println("Enter entry ID")
					scanner.Scan()
					entryId := scanner.Text()
					entry, err := journalInstance.GetEntry(strings.TrimSpace(entryId))
					if err != nil {
						fmt.Println(err)
						continue
					}
					fmt.Printf("ID: %s\n Title: %s\n Content: %s\n Created: %s\n Updated: %s\n\n", entry.ID, entry.Title, entry.Content, entry.Created, entry.Updated)

				case "delete":
					fmt.Println("Enter entry ID")
					scanner.Scan()
					entryId := scanner.Text()
					err := journalInstance.DeleteEntry(strings.TrimSpace(entryId))
					if err != nil {
						fmt.Println(err)
					}

				case "update":
					fmt.Println("Enter entry ID")
					scanner.Scan()
					entryId := scanner.Text()

					updatedEntry := map[string]string{
						"title":   "",
						"content": "",
					}
					for k, _ := range updatedEntry {
						fmt.Printf("Do you want to update the %v? Y/N\n", k)
						scanner.Scan()
						isUpdateField := strings.ToUpper(strings.TrimSpace(scanner.Text())) == "Y"
						if isUpdateField {
							fmt.Printf("Enter new %s\n", k)
							scanner.Scan()
							updatedEntry[k] = scanner.Text()
						}
					}
					entry, err := journalInstance.UpdateEntry(entryId, updatedEntry["title"], updatedEntry["content"])
					if err != nil {
						fmt.Println(err)
						continue
					}
					fmt.Printf("ID: %s\n Title: %s\n Content: %s\n Created: %s\n Updated: %s\n\n", entry.ID, entry.Title, entry.Content, entry.Created, entry.Updated)

				default:
					fmt.Println("Unknown command: " + input)
					fmt.Println("Available commands: create, list", "get", "delete", "update")

				}
			}
		}

	default:
		fmt.Println("Unknown command: " + command)
		fmt.Println("Available commands: create, list")

	}

}
