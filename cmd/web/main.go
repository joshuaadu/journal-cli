package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"journal/models"
	"journal/pkg/journal"
	"journal/pkg/storage"
	"journal/pkg/utils"
	"log"
	"net/http"
)

type EntryInput struct {
	Title, Content string
}

type PageData struct {
	Title         string
	Entries       []models.Entry
	Entry         models.Entry
	ShowCreateBtn bool
	Error         string
}

var journalIntance *journal.Journal

// Create a FuncMap with the custom time formatting function
var funcMap = template.FuncMap{
	"formatTime": utils.FormatTime,
}

func main() {
	//Initialize storage
	//db, err := storage.NewSQLiteStorage("journal.db")
	db, err := storage.NewMongoDBStorage("journal", "entries")
	if err != nil {
		log.Fatal("Failed to initialize storage: ", err)
	}

	journalIntance = journal.NewJournal(db)

	// Set up router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/test", TestHandler).Methods("GET")
	router.HandleFunc("/app", EntriesHandler).Methods("GET")
	router.HandleFunc("/app/entries/new", NewEntryPageHandler).Methods("GET")
	router.HandleFunc("/app/entries/new", PostNewEntryHandler).Methods("POST")
	router.HandleFunc("/app/entries/{id}", ViewEntryHandler).Methods("GET") // Get a specified entry by ID
	router.HandleFunc("/api/entries", ListEntries).Methods("GET")           // List all entries
	router.HandleFunc("/api/entries", CreateEntry).Methods("POST")          // Create a new entry
	router.HandleFunc("/api/entries/{id}", GetEntry).Methods("GET")         // Get a specified entry by ID
	router.HandleFunc("/api/entries/{id}", UpdateEntry).Methods("PUT")      //  // Update an entry by ID
	router.HandleFunc("/api/entries/{id}", DeleteEntry).Methods("DELETE")   // Delete an entry by ID

	// Start HTTP server
	port := ":8080"
	fmt.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func TestHandler(w http.ResponseWriter, r *http.Request) {

	templates := template.Must(template.ParseFiles("templates/pages/heloworld.html"))
	data := "Test"
	templates.Execute(w, data)
}

func ViewEntryHandler(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/layouts/base.html",
		"templates/pages/entry.html", "templates/partials/header.html"))
	data := PageData{
		Title:         "View Entry",
		ShowCreateBtn: true,
		//Error: "Method not allowed",
	}

	id := mux.Vars(r)["id"]
	entry, err := journalIntance.GetEntry(id)
	if err != nil {
		http.Redirect(w, r, "/app", http.StatusSeeOther)
	}
	data.Entry = entry

	err = templates.ExecuteTemplate(w, "base", data)

}

func PostNewEntryHandler(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/layouts/base.html",
		"templates/pages/new.html", "templates/partials/form.html", "templates/partials/header.html"))

	data := PageData{
		Title:         "Add New Entry",
		ShowCreateBtn: false,
		//Error: "Method not allowed",
	}
	if r.Method != http.MethodPost {
		data.Error = "Method not allowed"
		templates.ExecuteTemplate(w, "base", data)

		return
	}
	if err := r.ParseForm(); err != nil {
		data.Error = "Invalid input"
		fmt.Println(err)
		err = templates.ExecuteTemplate(w, "base", data)
		//log.Fatal("Create entry failed...:", err)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	if len(title) < 1 {
		data.Error = "Title is required!"
		err := templates.ExecuteTemplate(w, "base", data)
		fmt.Println(err)
		return
	}
	if len(content) < 1 {
		data.Error = "Content is required!"
		err := templates.ExecuteTemplate(w, "base", data)
		fmt.Println(err)
		return
	}
	_, err := journalIntance.CreateEntry(title, content)
	//fmt.Println("ran", entry)
	if err != nil {
		http.Error(w, "Failed to create entry", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/app", http.StatusSeeOther)
}

func NewEntryPageHandler(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/layouts/base.html",
		"templates/pages/new.html", "templates/partials/form.html", "templates/partials/header.html"))

	data := PageData{
		Title:         "Add New Entry",
		ShowCreateBtn: false,
	}
	templates.ExecuteTemplate(w, "base", data)
}

func EntriesHandler(w http.ResponseWriter, r *http.Request) {

	templates := template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/layouts/base.html",
		"templates/pages/entries.html", "templates/partials/header.html"))

	entries, _ := journalIntance.ListEntries()
	data := PageData{
		Title:         "Journal Entries",
		Entries:       entries,
		ShowCreateBtn: true,
	}
	//fmt.Printf("Home handler %+v\n", data)
	err := templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
