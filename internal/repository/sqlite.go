package repository

import "github.com/halalala222/GoBoilder/internal/template/repository"

var _ Template = &SQLiteRepositoryTemplate{}

// SQLiteRepositoryTemplate is the template for SQLite repository
type SQLiteRepositoryTemplate struct {
}

// Build returns the template for SQLite repository
func (t *SQLiteRepositoryTemplate) Build() []byte {
	return repository.SqliteTemplate
}
