package models

import "time"

// Entry represent a single journal entry
type Entry struct {
	ID      string    // Unique identifier for the entry
	Title   string    // Title of the journal entry
	Content string    // Content or body of the journal entry
	Created time.Time // Timestamp of when the entry was created
	Updated time.Time // Timestamp of when the entry was last updated
}

// UpdateEntry allows you to update the content and title of an existing entry.
func (entry *Entry) UpdateEntry(title, content string) {
	isUpdated := false
	if len(title) > 0 {
		entry.Title = title
		isUpdated = true
	}
	if len(content) > 0 {
		entry.Content = content
		isUpdated = true
	}
	if isUpdated {
		entry.Updated = time.Now()
	}
}
