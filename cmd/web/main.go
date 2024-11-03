package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"journal/pkg/journal"
	"journal/pkg/storage"
	"log"
	"net/http"
)

type EntryInput struct {
	Title, Content string
}

var journalIntance *journal.Journal

func main() {
	//Initialize storage
	db, err := storage.NewSQLiteStorage("journal.db")
	if err != nil {
		log.Fatal("Failed to initialize storage: ", err)
	}

	journalIntance = journal.NewJournal(db)

	// Set up router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/entries", ListEntries).Methods("GET")         // List all entries
	router.HandleFunc("/entries", CreateEntry).Methods("POST")        // Create a new entry
	router.HandleFunc("/entries/{id}", GetEntry).Methods("GET")       // Get a specified entry by ID
	router.HandleFunc("/entries/{id}", UpdateEntry).Methods("PUT")    //  // Update an entry by ID
	router.HandleFunc("/entries/{id}", DeleteEntry).Methods("DELETE") // Delete an entry by ID

	// Start HTTP server
	port := ":8080"
	fmt.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, router))

}

// GetEntry fetches a specific entry by ID
func GetEntry(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	entry, err := journalIntance.GetEntry(id)
	if err != nil {
		http.Error(w, "Entry not found", http.StatusNotFound)

	}
	err = json.NewEncoder(w).Encode(entry)
	if err != nil {
		log.Fatal("Sending get entry response failed:", err)
	}
}

// ListEntries list all entries
func ListEntries(w http.ResponseWriter, r *http.Request) {
	entries, err := journalIntance.ListEntries()

	if err != nil {
		http.Error(w, "Failed to fetch entries", http.StatusInternalServerError)
		log.Fatal("List entries failed:", err)
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(entries)
	if err != nil {
		log.Fatal("Sending entries response failed:", err)
	}

}

// CreateEntry creates a new journal entry
func CreateEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var entryInput EntryInput
	if err := json.NewDecoder(r.Body).Decode(&entryInput); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Fatal("Create entry failed:", err)
		return
	}

	entry, err := journalIntance.CreateEntry(entryInput.Title, entryInput.Content)
	if err != nil {
		http.Error(w, "Failed to create entry", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(entry)
	if err != nil {
		log.Fatal("Sending created entry response failed:", err)
	}

}

// UpdateEntry updates an entry by ID
func UpdateEntry(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var updateEntryInput EntryInput
	if err := json.NewDecoder(r.Body).Decode(&updateEntryInput); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	entry, err := journalIntance.UpdateEntry(id, updateEntryInput.Title, updateEntryInput.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(entry)
	if err != nil {
		log.Fatal("Sending updated entry response failed:", err)
	}
}

// DeleteEntry deletes an entry by ID
func DeleteEntry(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := journalIntance.DeleteEntry(id)
	if err != nil {
		http.Error(w, "Entry not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
