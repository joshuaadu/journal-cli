package journal

import (
	"errors"
)

// Journal holds a collection of entries.
type Journal struct {
	entries map[string]Entry // A map to store entries with their ID as the key
}

// NewJournal creates a new instance of Journal.
func NewJournal() *Journal {
	return &Journal{
		entries: make(map[string]Entry),
	}
}

// CreateEntry creates a new journal entry and adds it to the journal.
func (journal *Journal) CreateEntry(title, content string) Entry {
	entry := NewEntry(title, content)
	journal.entries[entry.ID] = entry
	return entry
}

// ListEntries returns all the entries in the journal
func (journal *Journal) ListEntries() []Entry {
	var entries []Entry
	for _, entry := range journal.entries {
		entries = append(entries, entry)
	}
	return entries
}

// UpdateEntry updates the title and content of an existing entry
func (journal *Journal) UpdateEntry(id, title, content string) (Entry, error) {
	entry, err := journal.GetEntry(id)
	if err != nil {
		return Entry{}, err
	}
	entry.UpdateEntry(title, content)
	journal.entries[id] = entry
	return entry, nil
}

// GetEntry retrieves a single entry by its ID.
func (journal *Journal) GetEntry(id string) (Entry, error) {
	entry, exists := journal.entries[id]
	if !exists {
		return Entry{}, errors.New("entry not found")
	}
	return entry, nil
}

// DeleteEntry removes  an entry from the journal by its ID.
func (journal *Journal) DeleteEntry(id string) error {
	if _, exists := journal.entries[id]; !exists {
		return errors.New("entry not found")
	}

	delete(journal.entries, id)
	return nil
}
