package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"journal/models"
)

type SQLiteStorage struct {
	DB *sql.DB
}

// NewSQLiteStorage initializes the SQLite database and returns a storage instance
func NewSQLiteStorage(dbFile string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	// Create journal_entries table if it doesn't exist
	query := `
    CREATE TABLE IF NOT EXISTS journal_entries (
        id TEXT PRIMARY KEY,
        title TEXT,
        content TEXT,
        created TIMESTAMP,
        updated TIMESTAMP
    );
    `
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{DB: db}, nil
}

// GetEntry loads a journal entry from the SQLite database
func (s *SQLiteStorage) GetEntry(id string) (models.Entry, error) {
	query := `
	SELECT id, title, content, created, updated FROM journal_entries WHERE id = ?
	`
	row := s.DB.QueryRow(query, id)
	var entry models.Entry
	err := row.Scan(&entry.ID, &entry.Title, &entry.Content, &entry.Created, &entry.Updated)
	if err != nil {
		return entry, err
	}
	return entry, nil
}

// LoadEntries loads journal entries from the SQLite database
func (s *SQLiteStorage) LoadEntries() ([]models.Entry, error) {
	query := `SELECT id, title, content, created, updated FROM journal_entries `
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []models.Entry
	for rows.Next() {
		var entry models.Entry
		err := rows.Scan(&entry.ID, &entry.Title, &entry.Content, &entry.Created, &entry.Updated)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// SaveEntries saves the journal entries to the SQLite database (insert or update)
func (s *SQLiteStorage) SaveEntries(entries []models.Entry) error {
	for _, entry := range entries {
		query := `
		INSERT INTO journal_entries (id, title, content, created, updated)
		VALUES (?,?,?,?,?)
		ON CONFLICT(id) DO UPDATE SET 
			title=excluded.title,
			content=excluded.content,
			updated=excluded.updated;
		`
		_, err := s.DB.Exec(query, entry.ID, entry.Title, entry.Created, entry.Created, entry.Updated)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateEntry creates a new journal entry in SQLite
func (s *SQLiteStorage) CreateEntry(entry models.Entry) error {
	query := `
	INSERT INTO journal_entries (id, title, content, created, updated)
	VALUES (?,?,?,?,?)
	`

	_, err := s.DB.Exec(query, entry.ID, entry.Title, entry.Content, entry.Created, entry.Updated)

	return err
}

// UpdateEntry updates an existing journal entry in SQLite
func (s *SQLiteStorage) UpdateEntry(entry models.Entry) error {
	query := `
	UPDATE journal_entries 
	SET title = ?, content = ?, updated = ?
	WHERE id = ?
	`
	_, err := s.DB.Exec(query, entry.Title, entry.Content, entry.Updated, entry.ID)
	return err
}

// DeleteEntry deletes a journal entry from SQLite by its ID
func (s *SQLiteStorage) DeleteEntry(id string) error {
	query := `
	DELETE FROM journal_entries WHERE id = ?
	`
	_, err := s.DB.Exec(query, id)
	return err
}
