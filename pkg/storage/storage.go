package storage

import (
	"journal/models"
)

// Storage interface defines methods for storing journal entries
type Storage interface {
	LoadEntries() ([]models.Entry, error)
	SaveEntries(entries []models.Entry) error
	CreateEntry(entry models.Entry) error
	UpdateEntry(entry models.Entry) error
	DeleteEntry(id string) error
	GetEntry(id string) (models.Entry, error)
}
