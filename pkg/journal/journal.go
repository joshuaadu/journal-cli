package journal

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
