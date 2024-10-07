package main

import (
	"bufio"
	"fmt"
	"journal/pkg/journal"
	"os"
	"strings"
)

func main() {
	// Create a new journal instance
	journalInstance := journal.NewJournal()

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
	case "interactive":
		// Create a scanner to read input from the terminal
		scanner := bufio.NewScanner(os.Stdin)
		var title, content string
		fmt.Println("You are in the Journal program in interactive mode")
		fmt.Println("usage: create | list")
		fmt.Println("Type 'exit' to quit.")
		for {
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
					fmt.Println("Enter journal title")
					scanner.Scan()
					fmt.Println("Enter journal content")
					title = scanner.Text()
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
				default:
					fmt.Println("Unknown command: " + input)
					fmt.Println("Available commands: create, list")

				}
			}
		}

	default:
		fmt.Println("Unknown command: " + command)
		fmt.Println("Available commands: create, list")

	}

}
