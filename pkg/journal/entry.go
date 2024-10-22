package journal

import (
	"journal/models"
	"journal/pkg/utils"
	"time"
)

// NewEntry creates a new journal entry with the current timestamp
func NewEntry(title, content string) models.Entry {
	return models.Entry{
		ID:      utils.GenerateID(),
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}
}
