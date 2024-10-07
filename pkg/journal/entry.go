package journal

import (
	"journal/pkg/utils" // Import the utils package
	"time"
)

// Entry represent a single journal entry
type Entry struct {
	ID      string    // Unique identifier for the entry
	Title   string    // Title of the journal entry
	Content string    // Content or body of the journal entry
	Created time.Time // Timestamp of when the entry was created
	Updated time.Time // Timestamp of when the entry was last updated
}

// NewEntry creates a new journal entry with the current timestamp
func NewEntry(title, content string) Entry {
	return Entry{
		ID:      utils.GenerateID(),
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}
}

// UpdateEntry allows you to update the content and title of an existing entry.
func (entry *Entry) UpdateEntry(title, content string) {
	entry.Title = title
	entry.Content = content
	entry.Updated = time.Now()
}
