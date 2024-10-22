package journal

import (
	"errors"
	"journal/models"
	"journal/pkg/storage"
)

// Journal holds a collection of entries.
type Journal struct {
	//entries map[string]Entry // A map to store entries with their ID as the key
	storage storage.Storage
}

// NewJournal creates a new instance of Journal.
func NewJournal(store storage.Storage) *Journal {
	return &Journal{
		//entries: make(map[string]Entry),
		storage: store,
	}
}

// CreateEntry creates a new journal entry and adds it to the journal.
func (journal *Journal) CreateEntry(title, content string) models.Entry {
	entry := NewEntry(title, content)
	//journal.entries[entry.ID] = entry
	journal.storage.CreateEntry(entry)
	return entry
}

// ListEntries returns all the entries in the journal
func (journal *Journal) ListEntries() []models.Entry {
	//var entries []Entry
	//for _, entry := range journal.entries {
	//	entries = append(entries, entry)
	//}
	entries, _ := journal.storage.LoadEntries()
	return entries
}

// UpdateEntry updates the title and content of an existing entry
func (journal *Journal) UpdateEntry(id, title, content string) (models.Entry, error) {
	entry, err := journal.GetEntry(id)
	if err != nil {
		return models.Entry{}, err
	}
	entry.UpdateEntry(title, content)
	//journal.entries[id] = entry
	journal.storage.UpdateEntry(entry)
	return entry, nil
}

// GetEntry retrieves a single entry by its ID.
func (journal *Journal) GetEntry(id string) (models.Entry, error) {
	//entry, exists := journal.entries[id]
	entry, err := journal.storage.GetEntry(id)
	if err != nil {
		return models.Entry{}, errors.New("entry not found")
	}
	return entry, nil
}

// DeleteEntry removes  an entry from the journal by its ID.
func (journal *Journal) DeleteEntry(id string) error {
	//if _, exists := journal.entries[id]; !exists {
	//	return errors.New("entry not found")
	//}
	err := journal.storage.DeleteEntry(id)
	if err != nil {
		return errors.New("entry not found")
	}
	//delete(journal.entries, id)
	return nil
}
